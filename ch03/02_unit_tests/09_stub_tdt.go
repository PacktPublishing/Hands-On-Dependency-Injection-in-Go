package unit_tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPersonNameStubs(t *testing.T) {
	// this value does not matter as the stub ignores it
	fakeID := 1

	scenarios := []struct {
		desc         string
		loaderStub   *PersonLoaderStub
		expectedName string
		expectErr    bool
	}{
		{
			desc: "happy path",
			loaderStub: &PersonLoaderStub{
				Person: &Person{Name: "Sophia"},
			},
			expectedName: "Sophia",
			expectErr:    false,
		},
		{
			desc: "input error",
			loaderStub: &PersonLoaderStub{
				Error: ErrNotFound,
			},
			expectedName: "",
			expectErr:    true,
		},
		{
			desc: "system error path",
			loaderStub: &PersonLoaderStub{
				Error: errors.New("something failed"),
			},
			expectedName: "",
			expectErr:    true,
		},
	}

	for _, scenario := range scenarios {
		result, resultErr := LoadPersonName(scenario.loaderStub, fakeID)

		assert.Equal(t, scenario.expectedName, result, scenario.desc)
		assert.Equal(t, scenario.expectErr, resultErr != nil, scenario.desc)
	}
}
