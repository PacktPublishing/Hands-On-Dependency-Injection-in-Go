package jit_injection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyLoadPersonLogicJIT(t *testing.T) {
	// setup the mock db
	mockDB := &mockDB{
		out: Person{Name: "Fred"},
	}

	// call the object we are testing
	testObj := MyLoadPersonLogicJIT{
		dataSource: mockDB,
	}
	result, resultErr := testObj.Load(123)

	// validate expectations
	assert.Equal(t, Person{Name: "Fred"}, result)
	assert.Nil(t, resultErr)
}
