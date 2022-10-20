// Package pkg is a folder containing libs for usage on main project
package pkg

import (
	"sync"
)

var lock = &sync.Mutex{}

var instance *Prometheus

// GetPrometheusInstance for singleton prometheus
func GetPrometheusInstance() *Prometheus {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		instance = new(Prometheus)
		instance.gaugeList = []typeGauge{}
		instance.counterList = []typeCounter{}
		instance.histogramList = []typeHistogram{}
	}
	return instance
}
