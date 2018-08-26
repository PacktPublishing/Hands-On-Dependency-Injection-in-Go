package logging

// Logger logs stuff
type Logger struct {
	Level int
}

// Error outputs a log at level ERROR
func (l *Logger) Error(message string, args ...interface{}) {
	// TODO: implement
}

// Warn outputs a log at level ERROR
func (l *Logger) Warn(message string, args ...interface{}) {
	// TODO: implement
}

// Info outputs a log at level ERROR
func (l *Logger) Info(message string, args ...interface{}) {
	// TODO: implement
}

// Debug outputs a log at level ERROR
func (l *Logger) Debug(message string, args ...interface{}) {
	// TODO: implement
}
