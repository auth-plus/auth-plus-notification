package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

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
			Name: "request_latency",
			Help: "Gauge request latency",
		})
		succeedCounter := prometheus.NewCounter(prometheus.CounterOpts{
			Name: "request_latency",
			Help: "Gauge request latency",
		})

		completionTime.Set(latency)
		if status >= 500 {
			errorCounter.Inc()
		} else {
			succeedCounter.Inc()
		}
	}
}
