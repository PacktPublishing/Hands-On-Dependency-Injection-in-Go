package unit_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Round the supplied number to the nearest integer
func Round(in float64) int {
	return 0
}

func TestRound_down(t *testing.T) {
	in := float64(1.1)
	expected := 1

	result := Round(in)
	assert.Equal(t, expected, result)
}

func TestRound_up(t *testing.T) {
	in := float64(3.7)
	expected := 4

	result := Round(in)
	assert.Equal(t, expected, result)
}

func TestRound_noChange(t *testing.T) {
	in := float64(6.0)
	expected := 6

	result := Round(in)
	assert.Equal(t, expected, result)
}
