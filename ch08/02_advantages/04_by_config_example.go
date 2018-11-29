package config_injection

// NewByConfigConstructor is the constructor for MyStruct
func NewByConfigConstructor(cfg MyConfig, url string, credentials string) *MyStruct {
	return &MyStruct{
		// code removed
	}
}

// MyConfig defines the config for MyStruct
type MyConfig interface {
	Logger() Logger
	Instrumentation() Instrumentation
	RateLimiter() RateLimiter
	Cache() Cache
}
