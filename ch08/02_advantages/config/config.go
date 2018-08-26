package config

import (
	"sync"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/02_advantages/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/02_advantages/stats"
)

// LoadFromFile loads the config from the supplied path
func LoadFromFile(path string) (*Config, error) {
	// TODO: implement
	return &Config{}, nil
}

// Config is the result of loading config from a file
type Config struct {
	// Log config
	LogLevel       int `json:"log_level"`
	logger         *logging.Logger
	loggerInitOnce sync.Once

	// Instrumentation config
	StatsDHostAndPort string `json:"stats_d_host_and_port"`
	stats             *stats.Collector
	statsInitOnce     sync.Once

	// Rate Limiter config
	RateLimiterMaxConcurrent int `json:"rate_limiter_max_concurrent"`
}

func (c *Config) Logger() *logging.Logger {
	c.loggerInitOnce.Do(func() {
		// use log level to create new logger
		c.logger = &logging.Logger{
			Level: c.LogLevel,
		}
	})

	return c.logger
}

func (c *Config) Stats() *stats.Collector {
	c.statsInitOnce.Do(func() {
		c.stats = &stats.Collector{
			HostAndPort: c.StatsDHostAndPort,
		}
	})

	return c.stats
}
