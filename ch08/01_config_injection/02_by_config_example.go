package config_injection

import (
	"time"
)

// NewByConfigConstructor is the constructor for MyStruct
func NewByConfigConstructor(cfg MyConfig, limiter RateLimiter, cache Cache) *MyStruct {
	return &MyStruct{
		// code removed
	}
}

// MyConfig defines the config for MyStruct
type MyConfig interface {
	Logger() Logger
	Instrumentation() Instrumentation
	Timeout() time.Duration
	Workers() int
}
