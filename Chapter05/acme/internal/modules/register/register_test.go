package register

import (
	"errors"
	"testing"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch05/acme/internal/modules/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterer_Do_happyPath(t *testing.T) {
	// monkey patch calls to the data package
	defer func(original func(in *data.Person) (int, error)) {
		// restore original
		saver = original
	}(saver)

	// replace method
	saver = func(in *data.Person) (int, error) {
		result := 888
		var resultErr error

		return result, resultErr
	}
	// end of monkey patch

	// inputs
	in := &data.Person{
		FullName: "Chang",
		Phone:    "11122233355",
		Currency: "CNY",
	}

	// call method
	registerer := &Registerer{}
	ID, err := registerer.Do(in)

	// validate expectations
	require.NoError(t, err)
	assert.Equal(t, 888, ID)
}

func TestRegisterer_Do_error(t *testing.T) {
	// monkey patch calls to the data package
	defer func(original func(in *data.Person) (int, error)) {
		// restore original
		saver = original
	}(saver)

	// replace method
	saver = func(in *data.Person) (int, error) {
		var result int
		resultErr := errors.New("failed to save")

		return result, resultErr
	}
	// end of monkey patch

	// inputs
	in := &data.Person{
		FullName: "Chang",
		Phone:    "11122233355",
		Currency: "CNY",
	}

	// call method
	registerer := &Registerer{}
	ID, err := registerer.Do(in)

	// validate expectations
	require.Error(t, err)
	assert.Equal(t, 0, ID)
}
