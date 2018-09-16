package jit_injection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLoadPersonLogic(t *testing.T) {
	// setup the mock db
	mockDB := &mockDB{
		out: Person{Name: "Fred"},
	}

	// call the object we are testing
	testObj := NewMyLoadPersonLogic(mockDB)
	result, resultErr := testObj.Load(123)

	// validate expectations
	assert.Equal(t, Person{Name: "Fred"}, result)
	assert.Nil(t, resultErr)
}

// mock implementation of DataSource
type mockDB struct {
	out    Person
	outErr error
}

// Load implements DataSource
func (m *mockDB) Load(ID int) (Person, error) {
	return m.out, m.outErr
}
