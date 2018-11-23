package list

import (
	"context"
	"errors"
	"testing"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLister_Do_happyPath(t *testing.T) {
	// monkey patch calls to the data package
	defer func(original func(_ context.Context, _ data.Config) ([]*data.Person, error)) {
		// restore original
		loader = original
	}(loader)

	// replace method
	loader = func(_ context.Context, _ data.Config) ([]*data.Person, error) {
		result := []*data.Person{
			{
				ID:       1234,
				FullName: "Sally",
			},
			{
				ID:       5678,
				FullName: "Jane",
			},
		}
		var resultErr error

		return result, resultErr
	}
	// end of monkey patch

	// call method
	lister := &Lister{}
	persons, err := lister.load()

	// validate expectations
	require.NoError(t, err)
	assert.Equal(t, 2, len(persons))
}

func TestLister_Do_noResults(t *testing.T) {
	// monkey patch calls to the data package
	defer func(original func(_ context.Context, _ data.Config) ([]*data.Person, error)) {
		// restore original
		loader = original
	}(loader)

	// replace method
	loader = func(_ context.Context, _ data.Config) ([]*data.Person, error) {
		var result []*data.Person
		resultErr := data.ErrNotFound

		return result, resultErr
	}
	// end of monkey patch

	// call method
	lister := &Lister{}
	persons, err := lister.load()

	// validate expectations
	require.Equal(t, errPeopleNotFound, err)
	assert.Equal(t, 0, len(persons))
}

func TestLister_Do_error(t *testing.T) {
	// monkey patch calls to the data package
	defer func(original func(_ context.Context, _ data.Config) ([]*data.Person, error)) {
		// restore original
		loader = original
	}(loader)

	// replace method
	loader = func(_ context.Context, _ data.Config) ([]*data.Person, error) {
		var result []*data.Person
		resultErr := errors.New("failed to load people")

		return result, resultErr
	}
	// end of monkey patch

	// call method
	lister := &Lister{}
	persons, err := lister.load()

	// validate expectations
	require.Error(t, err)
	assert.Equal(t, 0, len(persons))
}
