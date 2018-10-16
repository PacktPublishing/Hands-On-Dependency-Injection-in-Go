package improvements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogging(t *testing.T) {
	// build log recorder
	recorder := &LogRecorder{}

	// Call struct that uses a logger
	calculator := &Calculator{
		logger: recorder,
	}
	result := calculator.divide(10, 0)

	// validate expectations, including that the logger was called
	assert.Equal(t, 0, result)
	require.Equal(t, 1, len(recorder.Logs))
	assert.Equal(t, "cannot divide by 0", recorder.Logs[0])
}

type Calculator struct {
	logger Logger
}

func (c *Calculator) divide(dividend int, divisor int) int {
	if divisor == 0 {
		c.logger.Error("cannot divide by 0")
		return 0
	}

	return dividend / divisor
}

// Logger is our standard interface
type Logger interface {
	Error(message string, args ...interface{})
}

// LogRecorder implements Logger interface
type LogRecorder struct {
	Logs []string
}

func (l *LogRecorder) Error(message string, args ...interface{}) {
	// build log message
	logMessage := fmt.Sprintf(message, args...)

	// record log message
	l.Logs = append(l.Logs, logMessage)
}
