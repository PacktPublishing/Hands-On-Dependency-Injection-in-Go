package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLister_Do(t *testing.T) {
	// call method
	lister := &Lister{}
	persons, err := lister.load()

	// validate expectations
	require.NoError(t, err)
	assert.True(t, len(persons) >= 4)
}
