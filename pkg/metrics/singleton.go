// Package pkg is a folder containing libs for usage on main project
package pkg

import (
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var lock = &sync.RWMutex{}
var instance *Metrics

// GetMetricsInstance returns the singleton Metrics instance
func GetMetricsInstance() *Metrics {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = new(Metrics)
			instance.meter = otel.Meter("auth-plus-notification")
			instance.gauges = make(map[string]metric.Float64Histogram)
			instance.counters = make(map[string]metric.Float64Counter)
			instance.histograms = make(map[string]metric.Float64Histogram)
		}
	}
	return instance
}
