package middlewares

import (
	config "auth-plus-notification/config"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func Logger() gin.HandlerFunc {
	env := config.GetEnv()
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)

		completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "request_latency",
			Help: "Gauge request latency",
		})
		completionTime.SetToCurrentTime()
		if err := push.New(env.Prometheus.Url+":"+env.Prometheus.Port, "db_backup").
			Collector(completionTime).
			Grouping("db", "customers").
			Push(); err != nil {
			fmt.Println("Could not push completion time to Pushgateway:", err)
		}
	}
}
