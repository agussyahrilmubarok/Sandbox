package logging

import (
	"context"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

type ctxKeyRequestID struct{}

var RequestIDKey = ctxKeyRequestID{}

func RequestIDMiddleware(baseLogger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = xid.New().String()
			}

			c.Response().Header().Set("X-Request-ID", requestID)

			reqLogger := baseLogger.With().Str("request_id", requestID).Logger()

			ctx := context.WithValue(c.Request().Context(), RequestIDKey, requestID)
			ctx = reqLogger.WithContext(ctx)
			c.SetRequest(c.Request().WithContext(ctx))

			defer func() {
				if r := recover(); r != nil {
					stack := string(debug.Stack())
					reqLogger.Error().
						Interface("panic", r).
						Str("stack", stack).
						Msg("Recovered from panic")
					c.Response().WriteHeader(http.StatusInternalServerError)
					panic(r)
				}

				reqLogger.Info().
					Str("method", c.Request().Method).
					Str("url", c.Request().RequestURI).
					Int("status_code", c.Response().Status).
					Dur("elapsed", time.Since(start)).
					Msg("incoming request")
			}()

			return next(c)
		}
	}
}

func GetRequestID(ctx context.Context) string {
	if rid, ok := ctx.Value(RequestIDKey).(string); ok {
		return rid
	}
	return ""
}
