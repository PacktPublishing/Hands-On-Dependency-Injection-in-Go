package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/modules/get"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetHandler_ServeHTTP(t *testing.T) {
	scenarios := []struct {
		desc            string
		inRequest       func() *http.Request
		inModelMock     func() *MockGetModel
		expectedStatus  int
		expectedPayload string
	}{
		{
			desc: "happy path",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", "/person/1/", nil)
				require.NoError(t, err)

				// set values into request (required by the mux)
				return mux.SetURLVars(req, map[string]string{muxVarID: "1"})
			},
			inModelMock: func() *MockGetModel {
				output := &get.Person{
					ID:       1,
					FullName: "John",
					Phone:    "0123456789",
					Currency: "USD",
					Price:    100,
				}

				mockGetModel := &MockGetModel{}
				mockGetModel.On("Do", mock.Anything).Return(output, nil).Once()

				return mockGetModel
			},
			expectedStatus:  http.StatusOK,
			expectedPayload: `{"id":1,"name":"John","phone":"0123456789","currency":"USD","price":100}` + "\n",
		},
		{
			desc: "bad input (ID is invalid)",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", "/person/x/", nil)
				require.NoError(t, err)

				// set values into request (required by the mux)
				return mux.SetURLVars(req, map[string]string{muxVarID: "x"})
			},
			inModelMock: func() *MockGetModel {
				// expect the model not to be called
				mockRegisterModel := &MockGetModel{}
				return mockRegisterModel
			},
			expectedStatus:  http.StatusBadRequest,
			expectedPayload: ``,
		},
		{
			desc: "bad input (ID is missing)",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", "/person//", nil)
				require.NoError(t, err)

				// set values into request (required by the mux)
				return mux.SetURLVars(req, map[string]string{})
			},
			inModelMock: func() *MockGetModel {
				// expect the model not to be called
				mockRegisterModel := &MockGetModel{}
				return mockRegisterModel
			},
			expectedStatus:  http.StatusBadRequest,
			expectedPayload: ``,
		},
		{
			desc: "dependency fail",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", "/person/1/", nil)
				require.NoError(t, err)

				// set values into request (required by the mux)
				return mux.SetURLVars(req, map[string]string{muxVarID: "1"})
			},
			inModelMock: func() *MockGetModel {
				mockRegisterModel := &MockGetModel{}
				mockRegisterModel.On("Do", mock.Anything).Return(nil, errors.New("something failed")).Once()

				return mockRegisterModel
			},
			expectedStatus:  http.StatusNotFound,
			expectedPayload: ``,
		},
		{
			desc: "requested registration does not exist",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", "/person/1/", nil)
				require.NoError(t, err)

				// set values into request (required by the mux)
				return mux.SetURLVars(req, map[string]string{muxVarID: "1"})
			},
			inModelMock: func() *MockGetModel {
				mockRegisterModel := &MockGetModel{}
				mockRegisterModel.On("Do", mock.Anything).Return(nil, errors.New("person not found")).Once()

				return mockRegisterModel
			},
			expectedStatus:  http.StatusNotFound,
			expectedPayload: ``,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// define model layer mock
			mockGetModel := scenario.inModelMock()

			// build handler
			handler := NewGetHandler(&testConfig{}, mockGetModel)

			// perform request
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, scenario.inRequest())

			// validate outputs
			require.Equal(t, scenario.expectedStatus, response.Code, scenario.desc)

			payload, _ := ioutil.ReadAll(response.Body)
			assert.Equal(t, scenario.expectedPayload, string(payload), scenario.desc)
		})
	}
}

type testConfig struct {
}

func (t *testConfig) Logger() logging.Logger {
	return &logging.LoggerStdOut{}
}

func (*testConfig) BindAddress() string {
	return "0.0.0.0:0"
}
