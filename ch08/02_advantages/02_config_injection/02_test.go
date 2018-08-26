package config_injection

import (
	"testing"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/02_advantages/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/02_advantages/stats"
	"github.com/stretchr/testify/assert"
)

func TestConfigInjection(t *testing.T) {
	// build test config
	cfg := &TestConfig{}

	// build and use object
	obj := NewMyObject(cfg)
	result, resultErr := obj.Do()

	// validate
	assert.NotNil(t, result)
	assert.NoError(t, resultErr)
}

// Simple implementation of the Config interface
type TestConfig struct {
	logger *logging.Logger
	stats  *stats.Collector
}

func (t *TestConfig) Logger() *logging.Logger {
	return t.logger
}

func (t *TestConfig) Stats() *stats.Collector {
	return t.stats
}
