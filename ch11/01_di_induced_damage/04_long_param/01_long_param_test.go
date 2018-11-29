package long_param

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFancyFormatHandler(t *testing.T) {
	// inputs
	config := &stubConfig{}
	parser := &stubParser{}
	limiter := &stubRateLimiter{}
	loader := &stubLoader{}

	// create the handler
	fancyHandler := NewFancyFormatHandler(config, parser, limiter, loader)

	// call with fake HTTP request
	response := httptest.NewRecorder()
	request := &http.Request{}
	fancyHandler.ServeHTTP(response, request)

	// validate result
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "something fancy!", response.Body.String())
}

// define some mock implementations so that our test can run
type stubConfig struct{}

func (s *stubConfig) Logger() Logger {
	return &stubLogger{}
}

func (s *stubConfig) Instrumentation() Instrumentation {
	return &stubInstrumentation{}
}

type stubLogger struct{}

func (s *stubLogger) Error(message string, args ...interface{}) {
	// do nothing
}

func (s *stubLogger) Warn(message string, args ...interface{}) {
	// do nothing
}

func (s *stubLogger) Info(message string, args ...interface{}) {
	// do nothing
}

func (s *stubLogger) Debug(message string, args ...interface{}) {
	// do nothing
}

type stubInstrumentation struct{}

func (s *stubInstrumentation) Count(key string, value int) {
	// do nothing
}

func (s *stubInstrumentation) Duration(key string, start time.Time) {
	// do nothing
}

type stubParser struct{}

func (s *stubParser) Extract(req *http.Request) (int, error) {
	return 1, nil
}

type stubRateLimiter struct{}

func (s *stubRateLimiter) Acquire() {
	// do nothing
}

func (s *stubRateLimiter) Release() {
	// do nothing
}

type stubLoader struct{}

func (s *stubLoader) Load(ID int) ([]byte, error) {
	return []byte(`some data`), nil
}
