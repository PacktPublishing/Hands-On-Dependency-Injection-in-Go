package injected_config

import (
	"testing"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/02_advantages/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testConfigLocation = ""
)

func TestInjectedConfig(t *testing.T) {
	// load test config
	cfg, err := config.LoadFromFile(testConfigLocation)
	require.NoError(t, err)

	// build and use object
	obj := NewMyObject(cfg)
	result, resultErr := obj.Do()

	// validate
	assert.NotNil(t, result)
	assert.NoError(t, resultErr)
}
