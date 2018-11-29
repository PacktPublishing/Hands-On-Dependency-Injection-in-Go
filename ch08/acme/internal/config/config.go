package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
)

// DefaultEnvVar is the default environment variable the points to the config file
const DefaultEnvVar = "ACME_CONFIG"

// App is the application config
var App *Config

// Config defines the JSON format for the config file
type Config struct {
	// DSN is the data source name (format: https://github.com/go-sql-driver/mysql/#dsn-data-source-name)
	DSN string

	// Address is the IP address and port to bind this rest to
	Address string

	// BasePrice is the price of registration
	BasePrice float64

	// ExchangeRateBaseURL is the server and protocol part of the URL from which to load the exchange rate
	ExchangeRateBaseURL string

	// ExchangeRateAPIKey is the API for the exchange rate API
	ExchangeRateAPIKey string

	// environmental dependencies
	logger logging.Logger
}

// Logger returns a reference to the singleton logger
func (c *Config) Logger() logging.Logger {
	if c.logger == nil {
		c.logger = &logging.LoggerStdOut{}
	}

	return c.logger
}

// RegistrationBasePrice returns the base price for registrations
func (c *Config) RegistrationBasePrice() float64 {
	return c.BasePrice
}

// DataDSN returns the DSN
func (c *Config) DataDSN() string {
	return c.DSN
}

// ExchangeBaseURL returns the Base URL from which we can load exchange rates
func (c *Config) ExchangeBaseURL() string {
	return c.ExchangeRateBaseURL
}

// ExchangeAPIKey returns the DSN
func (c *Config) ExchangeAPIKey() string {
	return c.ExchangeRateAPIKey
}

// BindAddress returns the host and port this service should bind to
func (c *Config) BindAddress() string {
	return c.Address
}

// Load returns the config loaded from environment
func init() {
	filename, found := os.LookupEnv(DefaultEnvVar)
	if !found {
		logging.L.Error("failed to locate file specified by %s", DefaultEnvVar)
		return
	}

	_ = load(filename)
}

func load(filename string) error {
	App = &Config{}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		logging.L.Error("failed to read config file. err: %s", err)
		return err
	}

	err = json.Unmarshal(bytes, App)
	if err != nil {
		logging.L.Error("failed to parse config file. err : %s", err)
		return err
	}

	return nil
}
