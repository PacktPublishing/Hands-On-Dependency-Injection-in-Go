package advantages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	// monkey patch
	defer func(original *Rand) {
		// restore patch after use
		globalRand = original
	}(globalRand)

	// swap out for a predictable outcome
	globalRand = New(&stubSource{})
	// end monkey patch

	// call the function
	result := Int()
	assert.Equal(t, 234, result)
}

// this is a stubbed implementation of Source that returns a predictable value
type stubSource struct {
}

func (s *stubSource) Int63() int64 {
	return 234
}
