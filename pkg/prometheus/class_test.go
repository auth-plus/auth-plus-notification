package pkg

import (
	"fmt"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClassTestSuite struct {
	suite.Suite
}

func (suite *ClassTestSuite) Test_succeed_when_creating_counter() {
	mockData := MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	instance := new(Prometheus)
	instance.gaugeList = []typeGauge{}
	instance.counterList = []typeCounter{}
	instance.histogramList = []typeHistogram{}
	instance.CreateCounter(mockData.ID, mockData.Help)
	assert.Equal(suite.T(), instance.counterList[0].id, mockData.ID)
	assert.Equal(suite.T(), len(instance.gaugeList), 0)
	assert.Equal(suite.T(), len(instance.histogramList), 0)
}

func (suite *ClassTestSuite) Test_succeed_when_creating_gauge() {
	mockData := MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	instance := new(Prometheus)
	instance.gaugeList = []typeGauge{}
	instance.counterList = []typeCounter{}
	instance.histogramList = []typeHistogram{}
	instance.CreateGauge(mockData.ID, mockData.Help)
	assert.Equal(suite.T(), instance.gaugeList[0].id, mockData.ID)
	assert.Equal(suite.T(), len(instance.counterList), 0)
	assert.Equal(suite.T(), len(instance.histogramList), 0)
}

func (suite *ClassTestSuite) Test_succeed_when_creating_histogram() {
	mockData := MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	instance := new(Prometheus)
	instance.gaugeList = []typeGauge{}
	instance.counterList = []typeCounter{}
	instance.histogramList = []typeHistogram{}
	instance.CreateHistogram(mockData.ID, mockData.Help)
	assert.Equal(suite.T(), instance.histogramList[0].id, mockData.ID)
	assert.Equal(suite.T(), len(instance.counterList), 0)
	assert.Equal(suite.T(), len(instance.gaugeList), 0)
}

func (suite *ClassTestSuite) Test_succeed_when_using_counter() {
	mockData := MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	instance := new(Prometheus)
	counterMocked := new(CounterMocked)
	counterMocked.On("Inc").Return()
	instance.counterList = []typeCounter{{id: mockData.ID, counter: counterMocked}}
	instance.gaugeList = []typeGauge{}
	instance.histogramList = []typeHistogram{}
	instance.CounterIncrement(mockData.ID)
	counterMocked.MethodCalled("Inc")
}

func (suite *ClassTestSuite) Test_succeed_when_using_gauge() {
	mockData := MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	instance := new(Prometheus)
	gaugeMocked := new(GaugeMocked)
	gaugeMocked.On("Set", mockData.Number).Return()
	instance.counterList = []typeCounter{}
	instance.gaugeList = []typeGauge{{id: mockData.ID, gauge: gaugeMocked}}
	instance.histogramList = []typeHistogram{}
	instance.GaugeSet(mockData.ID, mockData.Number)
	gaugeMocked.MethodCalled("Set", mockData.Number)
}

func (suite *ClassTestSuite) Test_succeed_when_using_histogram() {
	mockData := MockedData{}
	errMock := faker.FakeData(&mockData)
	if errMock != nil {
		fmt.Println(errMock)
	}
	instance := new(Prometheus)
	histogramMocked := new(HistogramMocked)
	histogramMocked.On("Observe", mockData.Number).Return()
	instance.counterList = []typeCounter{}
	instance.gaugeList = []typeGauge{}
	instance.histogramList = []typeHistogram{{id: mockData.ID, histogram: histogramMocked}}
	instance.HistogramObserve(mockData.ID, mockData.Number)
	histogramMocked.MethodCalled("Observe", mockData.Number)
}

func TestClass(t *testing.T) {
	suite.Run(t, new(ClassTestSuite))
}
