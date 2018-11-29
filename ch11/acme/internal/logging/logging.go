package logging

import (
	"fmt"
)

// Logger is our standard interface
type Logger interface {
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}

// LoggerStdOut logs to std out
type LoggerStdOut struct{}

// Debug logs messages at DEBUG level
func (l LoggerStdOut) Debug(message string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+message, args...)
}

// Info logs messages at INFO level
func (l LoggerStdOut) Info(message string, args ...interface{}) {
	fmt.Printf("[INFO] "+message, args...)
}

// Warn logs messages at WARN level
func (l LoggerStdOut) Warn(message string, args ...interface{}) {
	fmt.Printf("[WARN] "+message, args...)
}

// Error logs messages at ERROR level
func (l LoggerStdOut) Error(message string, args ...interface{}) {
	fmt.Printf("[ERROR] "+message, args...)
}
