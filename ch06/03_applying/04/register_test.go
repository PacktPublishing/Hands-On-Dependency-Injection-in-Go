package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterHandler_ServeHTTP(t *testing.T) {
	scenarios := []struct {
		desc           string
		inRequest      func() *http.Request
		inModelMock    func() *MockRegisterModel
		expectedStatus int
		expectedHeader string
	}{
		// scenarios go here
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// define model layer mock
			mockRegisterModel := scenario.inModelMock()

			// build handler
			handler := &RegisterHandler{
				registerer: mockRegisterModel,
			}

			// perform request
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, scenario.inRequest())

			// validate outputs
			require.Equal(t, scenario.expectedStatus, response.Code)

			// call should output the location to the new person
			resultHeader := response.Header().Get("Location")
			assert.Equal(t, scenario.expectedHeader, resultHeader)

			// validate the mock was used as we expected
			assert.True(t, mockRegisterModel.AssertExpectations(t))
		})
	}
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
