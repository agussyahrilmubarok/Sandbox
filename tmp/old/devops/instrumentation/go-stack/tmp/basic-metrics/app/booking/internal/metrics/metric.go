package metrics

import (
	"example.com/booking/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMetricServer(cfg *config.Config) *echo.Echo {
	metricsServer := echo.New()
	metricsServer.HideBanner = true
	metricsServer.GET("/metrics/prometheus", func(c echo.Context) error {
		promhttp.Handler().ServeHTTP(c.Response(), c.Request())
		return nil
	})

	return metricsServer
}
