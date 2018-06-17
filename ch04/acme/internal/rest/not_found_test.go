package rest

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNotFoundHandler_ServeHTTP(t *testing.T) {
	// ensure the test always fails by giving it a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create and start a server
	// With out current implementation, we cannot test this handler without a full server as we need the mux.
	address, err := startServer(ctx)
	require.NoError(t, err)

	// build inputs
	response, err := http.Get("http://" + address + "/some-bad-address")

	// validate outputs
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, response.StatusCode)
}
