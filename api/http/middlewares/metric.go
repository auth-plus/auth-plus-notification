// Package middlewares contains all middleware for GIN
package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Metric is a middleware that gather metrics of system
func Metric() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request
		c.Next()

		// after request
		latency := float64(time.Since(t))

		// access the status we are sending
		status := c.Writer.Status()

		completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "request_latency",
			Help: "Gauge request latency",
		})
		errorCounter := prometheus.NewCounter(prometheus.CounterOpts{
			Name: "error_counter",
			Help: "Counter request 50X/40X",
		})
		succeedCounter := prometheus.NewCounter(prometheus.CounterOpts{
			Name: "succeed_counter",
			Help: "Counter request 20X",
		})

		completionTime.Set(latency)
		if status >= 500 {
			errorCounter.Inc()
		} else {
			succeedCounter.Inc()
		}
	}
}
