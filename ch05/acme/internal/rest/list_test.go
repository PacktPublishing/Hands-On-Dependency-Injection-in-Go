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

func TestListHandler_ServeHTTP(t *testing.T) {
	// ensure the test always fails by giving it a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create and start a server
	// With out current implementation, we cannot test this handler without a full server as we need the mux.
	address, err := startServer(ctx)
	require.NoError(t, err)

	// build inputs
	response, err := http.Get("http://" + address + "/person/list")

	// validate outputs
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, response.StatusCode)

	expectedPayload := []byte(`{"people":[{"id":1,"name":"John","phone":"0123456780"},{"id":2,"name":"Paul","phone":"0123456781"},{"id":3,"name":"George","phone":"0123456782"},{"id":4,"name":"Ringo","phone":"0123456783"}`)
	payload, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	// we have to use contains because other tests add more records
	assert.Contains(t, string(payload), string(expectedPayload))
}
