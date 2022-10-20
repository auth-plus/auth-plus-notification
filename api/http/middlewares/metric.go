// Package middlewares contains all middleware for GIN
package middlewares

import (
	pkg "auth-plus-notification/pkg/prometheus"
	"time"

	"github.com/gin-gonic/gin"
)

// Metric is a middleware that gather metrics of system
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		prom := pkg.GetPrometheusInstance()
		t := time.Now()
		c.Next()

		latency := float64(time.Since(t))
		status := c.Writer.Status()

		prom.GaugeSet("request_latency", latency)
		if status >= 500 {
			prom.CounterIncrement("error_counter")
		} else {
			prom.CounterIncrement("succeed_counter")
		}
	}
}

// MetricSetup is a function to register all metrics and instanciate the singleton
func MetricSetup() {
	prom := pkg.GetPrometheusInstance()
	prom.CreateGauge("request_latency", "Gauge request latency")
	prom.CreateCounter("error_counter", "Counter request 50X/40X")
	prom.CreateCounter("succeed_counter", "Counter request 20X")
}
