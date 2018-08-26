package register

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterer_Do_happyPath(t *testing.T) {
	// define context and therefore test timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// monkey patch calls to the data package
	defer func(original func(_ context.Context, _ data.Config, _ *data.Person) (int, error)) {
		// restore original
		saver = original
	}(saver)

	// replace method
	saver = func(_ context.Context, _ data.Config, _ *data.Person) (int, error) {
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
	registerer := &Registerer{
		cfg:       &testConfig{},
		exchanger: &stubExchanger{},
	}
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
	defer func(original func(_ context.Context, _ data.Config, _ *data.Person) (int, error)) {
		// restore original
		saver = original
	}(saver)

	// replace method
	saver = func(_ context.Context, _ data.Config, _ *data.Person) (int, error) {
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
	registerer := &Registerer{
		cfg:       &testConfig{},
		exchanger: &stubExchanger{},
	}
	ID, err := registerer.Do(ctx, in)

	// validate expectations
	require.Error(t, err)
	assert.Equal(t, 0, ID)
}

// Stub implementation of Config
type testConfig struct{}

// Logger implement Config
func (t *testConfig) Logger() logging.Logger {
	return &logging.LoggerStdOut{}
}

// RegistrationBasePrice implement Config
func (t *testConfig) RegistrationBasePrice() float64 {
	return 12.34
}

// DataDSN implements Config
func (t *testConfig) DataDSN() string {
	return ""
}

type stubExchanger struct{}

// Exchange implements Exchanger
func (s stubExchanger) Exchange(ctx context.Context, basePrice float64, currency string) (float64, error) {
	return 12.34, nil
}
