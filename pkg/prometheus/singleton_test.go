package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SingletonTestSuite struct {
	suite.Suite
}

func (suite *SingletonTestSuite) Test_succeed_when_instanciate_two_times() {
	instance := GetPrometheusInstance()
	instance2 := GetPrometheusInstance()
	assert.Equal(suite.T(), instance, instance2)
}

func TestSingleton(t *testing.T) {
	suite.Run(t, new(SingletonTestSuite))
}
