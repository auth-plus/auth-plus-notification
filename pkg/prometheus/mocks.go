// Package pkg is a folder containing libs for usage on main project
package pkg

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/mock"
)

// MockedData is a data generator randomly by faker
type MockedData struct {
	Help   string  `faker:"paragraph"`
	ID     string  `faker:"word"`
	Number float64 `faker:"amount"`
}

// CounterMocked is a mocked struct for prometheus.Counter
type CounterMocked struct {
	prometheus.Metric
	prometheus.Collector
	mock.Mock
}

// Inc a mocked function for counter
func (m *CounterMocked) Inc() {
	return
}

// Add a mocked function for counter
func (m *CounterMocked) Add(_ float64) {
	return
}

// GaugeMocked is a mocked struct for prometheus.Gauge
type GaugeMocked struct {
	prometheus.Metric
	prometheus.Collector
	mock.Mock
}

// Set a mocked function for gauge
func (m *GaugeMocked) Set(_ float64) {
	return
}

// Add a mocked function for gauge
func (m *GaugeMocked) Add(_ float64) {
	return
}

// Dec a mocked function for gauge
func (m *GaugeMocked) Dec() {
	return
}

// Inc a mocked function for gauge
func (m *GaugeMocked) Inc() {
	return
}

// SetToCurrentTime a mocked function for gauge
func (m *GaugeMocked) SetToCurrentTime() {
	return
}

// Sub a mocked function for gauge
func (m *GaugeMocked) Sub(_ float64) {
	return
}

// HistogramMocked is a mocked struct for prometheus.Histogram
type HistogramMocked struct {
	prometheus.Metric
	prometheus.Collector
	mock.Mock
}

// Observe a mocked function for histogram
func (m *HistogramMocked) Observe(_ float64) {
	return
}
