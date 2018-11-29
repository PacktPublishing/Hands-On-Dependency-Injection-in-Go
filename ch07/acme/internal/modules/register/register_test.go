package register

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch07/acme/internal/modules/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterer_Do_happyPath(t *testing.T) {
	// define context and therefore test timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// monkey patch calls to the data package
	defer func(original func(ctx context.Context, in *data.Person) (int, error)) {
		// restore original
		saver = original
	}(saver)

	// replace method
	saver = func(ctx context.Context, in *data.Person) (int, error) {
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
	ID, err := registerer.Do(ctx, in)

	// validate expectations
	require.NoError(t, err)
	assert.Equal(t, 888, ID)
}

func TestRegisterer_Do_error(t *testing.T) {
	// define context and therefore test timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// monkey patch calls to the data package
	defer func(original func(ctx context.Context, in *data.Person) (int, error)) {
		// restore original
		saver = original
	}(saver)

	// replace method
	saver = func(ctx context.Context, in *data.Person) (int, error) {
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
	ID, err := registerer.Do(ctx, in)

	// validate expectations
	require.Error(t, err)
	assert.Equal(t, 0, ID)
}
