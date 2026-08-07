package pkg

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type ClassTestSuite struct {
	suite.Suite
}

func (suite *ClassTestSuite) Test_succeed_when_creating_metrics() {
	instance := new(Metrics)
	instance.meter = otel.Meter("test")
	instance.gauges = make(map[string]metric.Float64Histogram)
	instance.counters = make(map[string]metric.Float64Counter)
	instance.histograms = make(map[string]metric.Float64Histogram)

	instance.CreateCounter("test_counter", "help")
	instance.CreateGauge("test_gauge", "help")
	instance.CreateHistogram("test_histogram", "help")
	
	suite.NotNil(instance.counters["test_counter"])
	suite.NotNil(instance.gauges["test_gauge"])
	suite.NotNil(instance.histograms["test_histogram"])
}

func TestClass(t *testing.T) {
	suite.Run(t, new(ClassTestSuite))
}
