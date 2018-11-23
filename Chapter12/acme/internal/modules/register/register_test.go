package register

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterer_Do_happyPath(t *testing.T) {
	// configure the mock saver
	mockResult := 888

	mockSaver := &mockMySaver{}
	mockSaver.On("Save", mock.Anything, mock.Anything).Return(mockResult, nil).Once()

	// define context and therefore test timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// inputs
	in := &Person{
		FullName: "Chang",
		Phone:    "11122233355",
		Currency: "CNY",
	}

	// call method
	registerer := &Registerer{
		cfg:       &testConfig{},
		exchanger: &stubExchanger{},
		data:      mockSaver,
	}
	ID, err := registerer.Do(ctx, in)

	// validate expectations
	require.NoError(t, err)
	assert.Equal(t, 888, ID)
	assert.True(t, mockSaver.AssertExpectations(t))
}

func TestRegisterer_Do_error(t *testing.T) {
	// configure the mock saver
	mockSaver := &mockMySaver{}
	mockSaver.On("Save", mock.Anything, mock.Anything).Return(0, errors.New("something failed")).Once()

	// define context and therefore test timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// inputs
	in := &Person{
		FullName: "Chang",
		Phone:    "11122233355",
		Currency: "CNY",
	}

	// call method
	registerer := &Registerer{
		cfg:       &testConfig{},
		exchanger: &stubExchanger{},
		data:      mockSaver,
	}
	ID, err := registerer.Do(ctx, in)

	// validate expectations
	require.Error(t, err)
	assert.Equal(t, 0, ID)
	assert.True(t, mockSaver.AssertExpectations(t))
}

func TestRegisterer_Do_exchangeError(t *testing.T) {
	// configure the mocks
	mockSaver := &mockMySaver{}
	mockExchanger := &MockExchanger{}
	mockExchanger.
		On("Exchange", mock.Anything, mock.Anything, mock.Anything).
		Return(0.0, errors.New("failed to load conversion")).
		Once()

	// define context and therefore test timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// inputs
	in := &Person{
		FullName: "Chang",
		Phone:    "11122233355",
		Currency: "CNY",
	}

	// call method
	registerer := &Registerer{
		cfg:       &testConfig{},
		exchanger: mockExchanger,
		data:      mockSaver,
	}
	ID, err := registerer.Do(ctx, in)

	// validate expectations
	require.Error(t, err)
	assert.Equal(t, 0, ID)
	assert.True(t, mockSaver.AssertExpectations(t))
	assert.True(t, mockExchanger.AssertExpectations(t))
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
