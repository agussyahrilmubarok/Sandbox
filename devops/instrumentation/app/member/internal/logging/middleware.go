package logging

import (
	"context"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
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

func TracingMiddleware(tracer trace.Tracer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(c.Request().Header))

			ctx, span := tracer.Start(ctx, c.Request().Method+" "+c.Path())
			defer span.End()

			traceID := span.SpanContext().TraceID().String()
			spanID := span.SpanContext().SpanID().String()
			log := zerolog.Ctx(ctx).With().
				Str("trace_id", traceID).
				Str("span_id", spanID).
				Logger()
			ctx = log.WithContext(ctx)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
