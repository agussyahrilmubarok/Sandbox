package tracing

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func Middleware(serviceName string) echo.MiddlewareFunc {
	tracer := otel.Tracer(serviceName)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			ctx := c.Request().Context()

			spanCtx, span := tracer.Start(ctx, c.Request().Method+" "+c.Path())
			defer span.End()

			c.SetRequest(c.Request().WithContext(spanCtx))

			err := next(c)

			status := c.Response().Status
			duration := time.Since(start)

			span.SetAttributes(
				attribute.String("http.method", c.Request().Method),
				attribute.String("http.path", c.Path()),
				attribute.Int("http.status_code", status),
				attribute.String("client.ip", c.RealIP()),
				attribute.Float64("http.duration_ms", float64(duration.Milliseconds())),
			)

			if err != nil {
				span.RecordError(err)
			}

			return err
		}
	}
}
