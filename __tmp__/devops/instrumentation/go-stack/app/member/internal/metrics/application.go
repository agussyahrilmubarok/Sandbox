package metrics

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// http_requests_total{job="member-app"}
	// http_requests_total{job="member-app", method="GET", status="200"}
	// rate(http_requests_total{job="member-app"}[5m])
	HTTPRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "status"},
	)

	// http_request_duration_seconds{job="member-app"}
	// avg(http_request_duration_seconds{job="member-app", method="GET", status="200"})
	// histogram_quantile(0.99, sum(rate(http_request_duration_seconds_bucket{job="member-app", method="GET", status="500"}[5m])) by (le))
	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of HTTP request durations in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "status"},
	)
)

func init() {
	prometheus.MustRegister(HTTPRequestCount)
	prometheus.MustRegister(RequestDuration)
}

func MetricAppMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		duration := time.Since(start).Seconds()

		HTTPRequestCount.WithLabelValues(c.Request().Method, http.StatusText(c.Response().Status)).Inc()
		RequestDuration.WithLabelValues(c.Request().Method, http.StatusText(c.Response().Status)).Observe(duration)

		return err
	}
}
