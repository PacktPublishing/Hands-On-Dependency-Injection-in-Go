package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	scenarios := []struct {
		desc           string
		in             string
		expectedConfig *Config
		expectError    bool
	}{
		{
			desc: "happy path",
			in:   "../../../../default-config.json",
			expectedConfig: &Config{
				DSN:                 "[insert your db config here]",
				Address:             "0.0.0.0:8080",
				BasePrice:           100.00,
				ExchangeRateBaseURL: "http://apilayer.net",
				ExchangeRateAPIKey:  "[insert your API key here]",
			},
			expectError: false,
		},
		{
			desc:           "invalid path",
			in:             "invalid.json",
			expectedConfig: &Config{},
			expectError:    true,
		},
	}

	for _, s := range scenarios {
		scenario := s
		t.Run(scenario.desc, func(t *testing.T) {
			resultErr := load(scenario.in)
			require.Equal(t, scenario.expectError, resultErr != nil, "err: %s", resultErr)
			assert.Equal(t, scenario.expectedConfig, App, scenario.desc)
		})
	}

}
