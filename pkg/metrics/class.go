// Package pkg is a folder containing libs for usage on main project
package pkg

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel/metric"
)

// Metrics is the main type for singleton
type Metrics struct {
	meter      metric.Meter
	gauges     map[string]metric.Float64Histogram
	counters   map[string]metric.Float64Counter
	histograms map[string]metric.Float64Histogram
}

// CreateGauge is a function using opentelemetry lib. Note: internally using Histogram since synchronous gauges are often misused for latencies.
func (p *Metrics) CreateGauge(id string, help string) {
	lock.Lock()
	defer lock.Unlock()
	
	if _, exists := p.gauges[id]; !exists {
		log.Println("registering gauge: ", id)
		instrument, err := p.meter.Float64Histogram(id, metric.WithDescription(help))
		if err != nil {
			log.Println("error creating gauge:", err)
			return
		}
		p.gauges[id] = instrument
	}
}

// CreateCounter is a function using opentelemetry lib
func (p *Metrics) CreateCounter(id string, help string) {
	lock.Lock()
	defer lock.Unlock()
	
	if _, exists := p.counters[id]; !exists {
		log.Println("registering counter: ", id)
		instrument, err := p.meter.Float64Counter(id, metric.WithDescription(help))
		if err != nil {
			log.Println("error creating counter:", err)
			return
		}
		p.counters[id] = instrument
	}
}

// CreateHistogram is a function using opentelemetry lib
func (p *Metrics) CreateHistogram(id string, help string) {
	lock.Lock()
	defer lock.Unlock()
	
	if _, exists := p.histograms[id]; !exists {
		log.Println("registering histogram: ", id)
		instrument, err := p.meter.Float64Histogram(id, metric.WithDescription(help))
		if err != nil {
			log.Println("error creating histogram:", err)
			return
		}
		p.histograms[id] = instrument
	}
}

// CounterIncrement is a function using opentelemetry lib
func (p *Metrics) CounterIncrement(id string) {
	lock.RLock()
	instrument, exists := p.counters[id]
	lock.RUnlock()
	
	if !exists {
		panic(fmt.Sprintf("Counter %s need to initalize first", id))
	}
	instrument.Add(context.Background(), 1)
}

// GaugeSet is a function using opentelemetry lib
func (p *Metrics) GaugeSet(id string, value float64) {
	lock.RLock()
	instrument, exists := p.gauges[id]
	lock.RUnlock()
	
	if !exists {
		panic(fmt.Sprintf("Gauge %s need to initalize first", id))
	}
	instrument.Record(context.Background(), value)
}

// HistogramObserve is a function using opentelemetry lib
func (p *Metrics) HistogramObserve(id string, value float64) {
	lock.RLock()
	instrument, exists := p.histograms[id]
	lock.RUnlock()
	
	if !exists {
		panic(fmt.Sprintf("Histogram %s need to initalize first", id))
	}
	instrument.Record(context.Background(), value)
}
