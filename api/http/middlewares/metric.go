// Package middlewares contains all middleware for GIN
package middlewares

import (
	pkg "auth-plus-notification/pkg/metrics"
	"time"

	"github.com/gin-gonic/gin"
)

// Metric is a middleware that gather metrics of system
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics := pkg.GetMetricsInstance()
		t := time.Now()
		c.Next()

		latency := float64(time.Since(t))
		status := c.Writer.Status()

		metrics.GaugeSet("request_latency", latency)
		if status >= 500 {
			metrics.CounterIncrement("error_counter")
		} else {
			metrics.CounterIncrement("succeed_counter")
		}
	}
}

// MetricSetup is a function to register all metrics and instanciate the singleton
func MetricSetup() {
	metrics := pkg.GetMetricsInstance()
	metrics.CreateGauge("request_latency", "Gauge request latency")
	metrics.CreateCounter("error_counter", "Counter request 50X/40X")
	metrics.CreateCounter("succeed_counter", "Counter request 20X")
}
