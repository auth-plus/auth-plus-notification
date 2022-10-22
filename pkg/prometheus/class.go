// Package pkg is a folder containing libs for usage on main project
package pkg

import (
	"fmt"
	"sort"

	"github.com/prometheus/client_golang/prometheus"
)

type typeGauge struct {
	id    string
	gauge prometheus.Gauge
}
type typeCounter struct {
	id      string
	counter prometheus.Counter
}
type typeHistogram struct {
	id        string
	histogram prometheus.Histogram
}

// Prometheus is the main type for singleton
type Prometheus struct {
	gaugeList     []typeGauge
	counterList   []typeCounter
	histogramList []typeHistogram
}

// CreateGauge is a function using prometheus lib
func (p *Prometheus) CreateGauge(id string, help string) {
	size := len(p.gaugeList)
	idx := sort.Search(size, func(i int) bool {
		return string(p.gaugeList[i].id) == id
	})
	print("CreateGauge", len(p.gaugeList), " --- ", idx)
	if idx == size {
		promGauge := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: id,
			Help: help,
		})
		gauge := typeGauge{
			id:    id,
			gauge: promGauge,
		}
		p.gaugeList = append(p.gaugeList, gauge)
		prometheus.MustRegister(promGauge)
	}
}

// CreateCounter is a function using prometheus lib
func (p *Prometheus) CreateCounter(id string, help string) {
	size := len(p.counterList)
	idx := sort.Search(size, func(i int) bool {
		return string(p.counterList[i].id) == id
	})
	if idx == size {
		promCounter := prometheus.NewCounter(prometheus.CounterOpts{
			Name: id,
			Help: help,
		})
		counter := typeCounter{
			id:      id,
			counter: promCounter,
		}
		p.counterList = append(p.counterList, counter)
		prometheus.MustRegister(promCounter)
	}
}

// CreateHistogram is a function using prometheus lib
func (p *Prometheus) CreateHistogram(id string, help string) {
	size := len(p.histogramList)
	idx := sort.Search(size, func(i int) bool {
		return string(p.histogramList[i].id) == id
	})
	if idx == size {
		promHistogram := prometheus.NewHistogram(prometheus.HistogramOpts{
			Name: id,
			Help: help,
		})
		histogram := typeHistogram{
			id:        id,
			histogram: promHistogram,
		}
		p.histogramList = append(p.histogramList, histogram)
		prometheus.MustRegister(promHistogram)
	}
}

// CounterIncrement is a function using prometheus lib
func (p *Prometheus) CounterIncrement(id string) {
	size := len(p.counterList)
	if size == 0 {
		panic(fmt.Sprintf("Counter %s need to initalize first", id))
	}
	idx := sort.Search(size, func(i int) bool {
		return string(p.counterList[i].id) == id
	})
	if idx != size {
		p.counterList[idx].counter.Inc()
	}
}

// GaugeSet is a function using prometheus lib
func (p *Prometheus) GaugeSet(id string, value float64) {
	size := len(p.gaugeList)
	if size == 0 {
		panic(fmt.Sprintf("Gauge %s need to initalize first", id))
	}
	idx := sort.Search(size, func(i int) bool {
		return string(p.gaugeList[i].id) == id
	})
	if idx != size {
		p.gaugeList[idx].gauge.Set(value)
	}
}
