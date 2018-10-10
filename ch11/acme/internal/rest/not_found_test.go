package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotFoundHandler_ServeHTTP(t *testing.T) {
	// build inputs
	response := httptest.NewRecorder()
	request := &http.Request{}

	// call handler
	notFoundHandler(response, request)

	// validate outputs
	require.Equal(t, http.StatusNotFound, response.Code)
}
