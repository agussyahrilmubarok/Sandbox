package metrics

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func PrometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start).Seconds()

		HTTPRequestCount.WithLabelValues(c.Request().Method, http.StatusText(c.Response().Status)).Inc()
		RequestDuration.WithLabelValues(c.Request().Method, http.StatusText(c.Response().Status)).Observe(duration)

		return err
	}
}
