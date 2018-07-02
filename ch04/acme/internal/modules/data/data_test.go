package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestData_happyPath(t *testing.T) {
	in := &Person{
		FullName: "Jake Blues",
		Phone:    "01234567890",
		Currency: "AUD",
		Price:    123.45,
	}

	// save
	resultID, err := Save(in)
	require.Nil(t, err)
	assert.True(t, resultID > 0)

	// load
	returned, err := Load(resultID)
	require.NoError(t, err)

	in.ID = resultID
	assert.Equal(t, in, returned)

	// load all
	all, err := LoadAll()
	require.NoError(t, err)
	assert.True(t, len(all) > 0)
}
