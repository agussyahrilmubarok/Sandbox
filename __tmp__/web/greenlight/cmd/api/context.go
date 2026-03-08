package main

import (
	"context"
	"net/http"

	"greenlight.id/internal/data"
)

// contextKey is a custom type used to prevent collisions between context keys in the application.
type contextKey string

// userContextKey is a constant key used to store and retrieve user data in the request context.
const userContextKey = contextKey("user")

// contextSetUser stores the user information in the request's context and returns a new request
// that includes the user data in the context. This allows user data to be accessed from the context
// throughout the request's lifecycle.
func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	// Create a new context with the user data attached to it.
	ctx := context.WithValue(r.Context(), userContextKey, user)
	// Return the updated request with the new context.
	return r.WithContext(ctx)
}

// contextGetUser retrieves the user information from the request's context.
// If the user data is not found, it will panic. This function assumes that
// the user data has already been set in the context earlier in the request lifecycle.
func (app *application) contextGetUser(r *http.Request) *data.User {
	// Extract the user data from the request's context using the userContextKey.
	user, ok := r.Context().Value(userContextKey).(*data.User)
	// If the user data is missing from the context, raise a panic.
	if !ok {
		panic("missing user value in request context")
	}
	// Return the user data.
	return user
}
