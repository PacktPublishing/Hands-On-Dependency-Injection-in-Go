package config_injection

import (
	"time"
)

// NewLongConstructor is the constructor for MyStruct
func NewLongConstructor(logger Logger, stats Instrumentation, limiter RateLimiter, cache Cache, url string, credentials string) *MyStruct {
	return &MyStruct{
		// code removed
	}
}

// MyStruct does something fantastic
type MyStruct struct {
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

// RateLimiter limits how many concurrent requests we can make or process
type RateLimiter interface {
	Acquire()
	Release()
}

// Cache will store/retrieve data in a fast way
type Cache interface {
	Store(key string, data []byte)
	Get(key string) ([]byte, error)
}
