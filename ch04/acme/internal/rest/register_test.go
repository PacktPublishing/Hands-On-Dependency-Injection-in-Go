package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterHandler_ServeHTTP(t *testing.T) {
	// ensure the test always fails by giving it a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create and start a server
	// With out current implementation, we cannot test this handler without a full server as we need the mux.
	address, err := startServer(ctx)
	require.NoError(t, err)

	// build inputs
	validRequest := buildValidRequest()
	response, err := http.Post("http://"+address+"/person/register", "application/json", validRequest)

	// validate outputs
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, response.StatusCode)
	defer response.Body.Close()

	// call should output the location to the new person
	headerLocation := response.Header.Get("Location")
	assert.Contains(t, headerLocation, "/person/")
}

func buildValidRequest() io.Reader {
	requestData := &registerRequest{
		FullName: "Joan Smith",
		Currency: "AUD",
		Phone:    "01234567890",
	}

	data, _ := json.Marshal(requestData)
	return bytes.NewBuffer(data)
}
