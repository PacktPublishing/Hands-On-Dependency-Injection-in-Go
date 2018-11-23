package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplements(t *testing.T) {
	assert.Implements(t, (*Database)(nil), &DatabaseImpl{})
	assert.Implements(t, (*Row)(nil), &RowImpl{})
	assert.Implements(t, (*Rows)(nil), &RowsImpl{})
	assert.Implements(t, (*Result)(nil), &ResultImpl{})
}
