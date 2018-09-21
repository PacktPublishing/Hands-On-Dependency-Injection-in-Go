package rest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/acme/internal/modules/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListHandler_ServeHTTP(t *testing.T) {
	scenarios := []struct {
		desc            string
		inRequest       func() *http.Request
		inModelMock     func() *MockListModel
		expectedStatus  int
		expectedPayload string
	}{
		{
			desc: "happy path",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", "/person/list", nil)
				require.NoError(t, err)

				return req
			},
			inModelMock: func() *MockListModel {
				output := []*data.Person{
					{
						ID:       1,
						FullName: "John",
						Phone:    "0123456789",
					},
					{
						ID:       2,
						FullName: "Paul",
						Phone:    "0123456781",
					},
					{
						ID:       3,
						FullName: "George",
						Phone:    "0123456782",
					},
					{
						ID:       1,
						FullName: "Ringo",
						Phone:    "0123456783",
					},
				}

				mockListModel := &MockListModel{}
				mockListModel.On("Do", mock.Anything).Return(output, nil).Once()

				return mockListModel
			},
			expectedStatus:  http.StatusOK,
			expectedPayload: `{"people":[{"id":1,"name":"John","phone":"0123456789"},{"id":2,"name":"Paul","phone":"0123456781"},{"id":3,"name":"George","phone":"0123456782"},{"id":1,"name":"Ringo","phone":"0123456783"}]}` + "\n",
		},
		{
			desc: "dependency failure",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", "/person/list", nil)
				require.NoError(t, err)

				return req
			},
			inModelMock: func() *MockListModel {
				mockListModel := &MockListModel{}
				mockListModel.On("Do", mock.Anything).Return(nil, errors.New("something failed")).Once()

				return mockListModel
			},
			expectedStatus:  http.StatusNotFound,
			expectedPayload: ``,
		},
		{
			desc: "no data",
			inRequest: func() *http.Request {
				req, err := http.NewRequest("GET", "/person/list", nil)
				require.NoError(t, err)

				return req
			},
			inModelMock: func() *MockListModel {
				// no data
				output := []*data.Person{}

				mockListModel := &MockListModel{}
				mockListModel.On("Do", mock.Anything).Return(output, nil).Once()

				return mockListModel
			},
			expectedStatus:  http.StatusOK,
			expectedPayload: `{"people":[]}` + "\n",
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			// define model layer mock
			mockListModel := scenario.inModelMock()

			// build handler
			handler := NewListHandler(mockListModel)

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
