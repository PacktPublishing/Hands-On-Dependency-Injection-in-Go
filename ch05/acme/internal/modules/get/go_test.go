package get

import (
	"errors"
	"testing"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch05/acme/internal/modules/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetter_Do_happyPath(t *testing.T) {
	// inputs
	ID := 1234

	// monkey patch calls to the data package
	defer func(original func(ID int) (*data.Person, error)) {
		// restore original
		loader = original
	}(loader)

	// replace method
	loader = func(ID int) (*data.Person, error) {
		result := &data.Person{
			ID:       1234,
			FullName: "Doug",
		}
		var resultErr error

		return result, resultErr
	}
	// end of monkey patch

	// call method
	getter := &Getter{}
	person, err := getter.Do(ID)

	// validate expectations
	require.NoError(t, err)
	assert.Equal(t, ID, person.ID)
	assert.Equal(t, "Doug", person.FullName)
}

func TestGetter_Do_noSuchPerson(t *testing.T) {
	// inputs
	ID := 5678

	// monkey patch calls to the data package
	defer func(original func(ID int) (*data.Person, error)) {
		// restore original
		loader = original
	}(loader)

	// replace method
	loader = func(ID int) (*data.Person, error) {
		var result *data.Person
		resultErr := data.ErrNotFound

		return result, resultErr
	}
	// end of monkey patch

	// call method
	getter := &Getter{}
	person, err := getter.Do(ID)

	// validate expectations
	require.Equal(t, errPersonNotFound, err)
	assert.Nil(t, person)
}

func TestGetter_Do_error(t *testing.T) {
	// inputs
	ID := 1234

	// monkey patch calls to the data package
	defer func(original func(ID int) (*data.Person, error)) {
		// restore original
		loader = original
	}(loader)

	// replace method
	loader = func(ID int) (*data.Person, error) {
		var result *data.Person
		resultErr := errors.New("failed to load person")

		return result, resultErr
	}
	// end of monkey patch

	// call method
	getter := &Getter{}
	person, err := getter.Do(ID)

	// validate expectations
	require.Error(t, err)
	assert.Nil(t, person)
}
