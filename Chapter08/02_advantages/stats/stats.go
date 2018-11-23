package stats

import (
	"time"
)

// Collector collects and forwards stats
type Collector struct {
	HostAndPort string
}

// Count will record an event
func (c *Collector) Count(key string, value int) {
	// TODO: implement
}

// Count will record the duration of an event
func (c *Collector) Duration(key string, start time.Time) {
	// TODO: implement
}
