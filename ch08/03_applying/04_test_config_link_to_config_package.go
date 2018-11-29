// +build do-not-build

package applying

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
)

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

// ExchangeBaseURL implements Config
func (t *testConfig) ExchangeBaseURL() string {
	return config.App.ExchangeRateBaseURL
}

// ExchangeAPIKey implements Config
func (t *testConfig) ExchangeAPIKey() string {
	return config.App.ExchangeRateAPIKey
}
