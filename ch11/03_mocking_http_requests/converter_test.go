package mocking_http_reques

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestExchange_invalidResponse(t *testing.T) {
	// build response
	response := httptest.NewRecorder()
	_, err := response.WriteString(`invalid payload`)
	require.NoError(t, err)

	// configure mock
	mockRequester := &mockRequester{}
	mockRequester.On("doRequest", mock.Anything, mock.Anything).Return(response.Result(), nil).Once()

	// inputs
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	basePrice := 12.34
	currency := "AUD"

	// perform call
	converter := &Converter{
		requester: mockRequester,
		cfg:       &testConfig{},
	}
	result, resultErr := converter.Exchange(ctx, basePrice, currency)

	// validate response
	assert.Equal(t, float64(0), result)
	assert.Error(t, resultErr)
	assert.True(t, mockRequester.AssertExpectations(t))
}

// stub config that returns known values
type testConfig struct {
}

func (t *testConfig) Logger() Logger {
	return &stubLogger{}
}

func (t *testConfig) ExchangeBaseURL() string {
	return "http://www.example.com"
}

func (t *testConfig) ExchangeAPIKey() string {
	return "foo"
}
