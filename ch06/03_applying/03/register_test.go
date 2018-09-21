package rest

import (
	"net/http"
	"testing"
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
			// test goes here
		})
	}
}
