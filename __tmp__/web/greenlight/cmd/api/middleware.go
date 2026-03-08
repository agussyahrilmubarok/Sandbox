package main

import (
	"errors"
	"expvar"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/felixge/httpsnoop"
	"github.com/pascaldekloe/jwt"
	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
	"greenlight.id/internal/data"
)

// recoverPanic is a middleware that recovers from panics in the HTTP handlers.
// It logs the error and responds with a server error.
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close") // Close the connection on error
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r) // Call the next handler
	})
}

// rateLimit is a middleware that limits the number of requests from each client
// based on their IP address. It implements a rate-limiting mechanism.
func (app *application) rateLimit(next http.Handler) http.Handler {
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client) // Track clients and their request limits
	)

	// Periodically clean up old clients from the map.
	go func() {
		for {
			time.Sleep(time.Minute) // Cleanup interval
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip) // Remove inactive clients
				}
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.config.limiter.enabled {
			ip := realip.FromRequest(r) // Get the client's IP address
			mu.Lock()
			if _, found := clients[ip]; !found {
				// Create a new limiter for new clients
				clients[ip] = &client{
					limiter: rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst),
				}
			}
			clients[ip].lastSeen = time.Now() // Update last seen time
			if !clients[ip].limiter.Allow() { // Check if request is allowed
				mu.Unlock()
				app.rateLimitExceededResponse(w, r) // Respond with rate limit exceeded
				return
			}
			mu.Unlock()
		}
		next.ServeHTTP(w, r) // Call the next handler
	})
}

// authenticate is a middleware that checks for a valid JWT token in the Authorization header.
// It retrieves the user associated with the token and sets it in the request context.
func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization") // Handle CORS preflight for Authorization header

		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			// No token provided; treat user as anonymous
			r = app.contextSetUser(r, data.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		token := headerParts[1]
		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret))
		if err != nil {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		// Validate the JWT claims
		if !claims.Valid(time.Now()) || claims.Issuer != "greenlight.id" || !claims.AcceptAudience("greenlight.id") {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		user, err := app.models.Users.Get(userID) // Retrieve user from the database
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.invalidAuthenticationTokenResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}
		// Add the user record to the request context
		r = app.contextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}

// requireActivatedUser checks if the user account is activated before proceeding.
// If not activated, it responds with an inactive account message.
func (app *application) requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)
		if !user.Activated {
			app.inactiveAccountResponse(w, r)
			return
		}
		next.ServeHTTP(w, r) // Proceed to the next handler
	})
	return app.requireAuthenticatedUser(fn) // Ensure user is authenticated
}

// requireAuthenticatedUser checks if the user is authenticated (not anonymous).
// If not, it responds with a required authentication message.
func (app *application) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)
		if user.IsAnonymous() {
			app.authenticationRequiredResponse(w, r)
			return
		}
		next.ServeHTTP(w, r) // Proceed to the next handler
	})
}

// requirePermission checks if the authenticated user has the required permission.
// If not, it responds with a not permitted message.
func (app *application) requirePermission(code string, next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)

		permissions, err := app.models.Permissions.GetAllForUser(user.ID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		if !permissions.Include(code) {
			app.notPermittedResponse(w, r)
			return
		}
		next.ServeHTTP(w, r) // Proceed to the next handler
	}

	return app.requireActivatedUser(fn) // Ensure user is activated
}

// enableCORS is a middleware that adds CORS headers to the response
// to allow cross-origin requests from trusted origins.
func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")

		origin := r.Header.Get("Origin")

		if origin != "" {
			// Check if the origin is trusted
			for i := range app.config.cors.trustedOrigins {
				if origin == app.config.cors.trustedOrigins[i] {
					w.Header().Set("Access-Control-Allow-Origin", origin)

					// Handle preflight request for CORS
					if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
						w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PUT, PATCH, DELETE")
						w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
						w.WriteHeader(http.StatusOK)
						return
					}
					break
				}
			}
		}

		next.ServeHTTP(w, r) // Proceed to the next handler
	})
}

// metrics is a middleware that tracks and exposes application metrics using expvar.
// It captures the total number of requests, responses, processing time, and response status counts.
func (app *application) metrics(next http.Handler) http.Handler {
	totalRequestsReceived := expvar.NewInt("total_requests_received")
	totalResponsesSent := expvar.NewInt("total_responses_sent")
	totalProcessingTimeMicroseconds := expvar.NewInt("total_processing_time_μs")

	totalResponsesSentByStatus := expvar.NewMap("total_responses_sent_by_status")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		totalRequestsReceived.Add(1) // Increment the total requests counter

		metrics := httpsnoop.CaptureMetrics(next, w, r) // Capture metrics for the response

		totalResponsesSent.Add(1)                                            // Increment the total responses counter
		totalProcessingTimeMicroseconds.Add(metrics.Duration.Microseconds()) // Add processing time
		totalResponsesSentByStatus.Add(strconv.Itoa(metrics.Code), 1)        // Increment status response count
	})
}
