package logging

import (
	"fmt"
)

// Debug logs messages at DEBUG level
func Debug(message string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+message, args...)
}

// Info logs messages at INFO level
func Info(message string, args ...interface{}) {
	fmt.Printf("[INFO] "+message, args...)
}

// Warn logs messages at WARN level
func Warn(message string, args ...interface{}) {
	fmt.Printf("[WARN] "+message, args...)
}

// Error logs messages at ERROR level
func Error(message string, args ...interface{}) {
	fmt.Printf("[ERROR] "+message, args...)
}
