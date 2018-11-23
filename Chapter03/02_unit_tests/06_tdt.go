package unit_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	scenarios := []struct {
		desc     string
		in       float64
		expected int
	}{
		{
			desc:     "round down",
			in:       1.1,
			expected: 1,
		},
		{
			desc:     "round up",
			in:       3.7,
			expected: 4,
		},
		{
			desc:     "unchanged",
			in:       6.0,
			expected: 6,
		},
	}

	for _, scenario := range scenarios {
		in := float64(scenario.in)

		result := Round(in)
		assert.Equal(t, scenario.expected, result)
	}
}
