// +build external

package exchange

import (
	"context"
	"testing"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExternalBoundaryTest(t *testing.T) {
	// define the config
	cfg, err := config.Load()
	require.NoError(t, err)

	// create a converter to test
	converter := NewConverter(cfg)

	// fetch from the server
	response, err := converter.loadRateFromServer(context.Background(), "AUD")
	require.NotNil(t, response)
	require.NoError(t, err)

	// parse the response
	resultRate, err := converter.extractRate(response, "AUD")
	require.NoError(t, err)

	// validate the result
	assert.True(t, resultRate > 0)
}
