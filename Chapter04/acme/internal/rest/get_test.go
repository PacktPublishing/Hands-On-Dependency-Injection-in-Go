package rest

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHandler_ServeHTTP(t *testing.T) {
	// ensure the test always fails by giving it a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create and start a server
	// With out current implementation, we cannot test this handler without a full server as we need the mux.
	address, err := startServer(ctx)
	require.NoError(t, err)

	// build inputs
	response, err := http.Get("http://" + address + "/person/1/")

	// validate outputs
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	expectedPayload := []byte(`{"id":1,"name":"John","phone":"0123456780","currency":"USD","price":100}` + "\n")
	payload, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	assert.Equal(t, expectedPayload, payload)
}
