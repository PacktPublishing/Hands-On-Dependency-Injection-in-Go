package config_injection

import (
	"fmt"
	"time"
)

func Usage() {
	cfg := &fakeConfig{}

	myFetcher := NewFetcher(cfg, cfg.URL(), cfg.Timeout())

	// do something with the object so the compiler does not complain
	fmt.Printf("%#v", myFetcher)
}

type FetcherConfig interface {
	Logger() Logger
	Instrumentation() Instrumentation
}

func NewFetcher(cfg FetcherConfig, url string, timeout time.Duration) *MyObject {
	return nil
}

type MyObject struct{}

// fake implementation of the FetcherConfig interface
type fakeConfig struct{}

// Logger implements FetcherConfig
func (f *fakeConfig) Logger() Logger {
	return nil
}

// Instrumentation implements FetcherConfig
func (f *fakeConfig) Instrumentation() Instrumentation {
	return nil
}

func (f *fakeConfig) URL() string {
	return ""
}

func (f *fakeConfig) Timeout() time.Duration {
	return time.Duration(0)
}
