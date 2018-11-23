package config_injection

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/02_advantages/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/02_advantages/stats"
)

func NewMyObject(cfg Config) *MyObject {
	return &MyObject{
		cfg: cfg,
	}
}

type Config interface {
	Logger() *logging.Logger
	Stats() *stats.Collector
}

type MyObject struct {
	cfg Config
}

func (m *MyObject) Do() (interface{}, error) {
	m.cfg.Logger().Error("not implemented")
	return struct{}{}, nil
}
