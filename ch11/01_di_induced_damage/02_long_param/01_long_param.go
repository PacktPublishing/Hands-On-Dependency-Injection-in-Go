package long_param

import (
	"net/http"
	"time"
)

func NewMyHandler(logger Logger, stats Instrumentation,
	parser Parser, formatter Formatter,
	limiter RateLimiter,
	loader Loader) *MyHandler {

	return &MyHandler{
		// code removed
	}
}

// MyHandler does something fantastic
type MyHandler struct {
	// code removed
}

func (m *MyHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// code removed
}

// Logger logs stuff
type Logger interface {
	Error(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Info(message string, args ...interface{})
	Debug(message string, args ...interface{})
}

// Instrumentation records the performances and events
type Instrumentation interface {
	Count(key string, value int)
	Duration(key string, start time.Time)
}

// Parse will extract details from the request
type Parser interface {
	Extract(req *http.Request) (int, error)
}

// Formatter will build the output
type Formatter interface {
	Format(resp http.ResponseWriter, data []byte) error
}

// RateLimiter limits how many concurrent requests we can make or process
type RateLimiter interface {
	Acquire()
	Release()
}

// Loader is responsible for loading the data
type Loader interface {
	Load(ID int) ([]byte, error)
}
