package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		{
			desc: "Happy Path",
			inRequest: func() *http.Request {
				validRequest := buildValidRegisterRequest()
				request, err := http.NewRequest("POST", "/person/register", validRequest)
				require.NoError(t, err)

				return request
			},
			inModelMock: func() *MockRegisterModel {
				// valid downstream configuration
				resultID := 1234
				var resultErr error

				mockRegisterModel := &MockRegisterModel{}
				mockRegisterModel.On("Do", mock.Anything, mock.Anything).Return(resultID, resultErr).Once()

				return mockRegisterModel
			},
			expectedStatus: http.StatusCreated,
			expectedHeader: "/person/1234/",
		},
		{
			desc: "Bad Input / User Error",
			inRequest: func() *http.Request {
				invalidRequest := bytes.NewBufferString(`this is not valid JSON`)
				request, err := http.NewRequest("POST", "/person/register", invalidRequest)
				require.NoError(t, err)

				return request
			},
			inModelMock: func() *MockRegisterModel {
				// Dependency should not be called
				mockRegisterModel := &MockRegisterModel{}
				return mockRegisterModel
			},
			expectedStatus: http.StatusBadRequest,
			expectedHeader: "",
		},
		{
			desc: "Dependency Failure",
			inRequest: func() *http.Request {
				validRequest := buildValidRegisterRequest()
				request, err := http.NewRequest("POST", "/person/register", validRequest)
				require.NoError(t, err)

				return request
			},
			inModelMock: func() *MockRegisterModel {
				// call to the dependency failed
				resultErr := errors.New("something failed")

				mockRegisterModel := &MockRegisterModel{}
				mockRegisterModel.On("Do", mock.Anything, mock.Anything).Return(0, resultErr).Once()

				return mockRegisterModel
			},
			expectedStatus: http.StatusBadRequest,
			expectedHeader: "",
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// define model layer mock
			mockRegisterModel := scenario.inModelMock()

			// build handler
			handler := NewRegisterHandler(mockRegisterModel)

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

func buildValidRegisterRequest() io.Reader {
	requestData := &registerRequest{
		FullName: "Joan Smith",
		Currency: "AUD",
		Phone:    "01234567890",
	}

	data, _ := json.Marshal(requestData)
	return bytes.NewBuffer(data)
}
