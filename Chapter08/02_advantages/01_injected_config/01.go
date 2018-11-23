package injected_config

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/02_advantages/config"
)

func NewMyObject(cfg *config.Config) *MyObject {
	return &MyObject{
		cfg: cfg,
	}
}

type MyObject struct {
	cfg *config.Config
}

func (m *MyObject) Do() (interface{}, error) {
	m.cfg.Logger().Error("not implemented")
	return struct{}{}, nil
}
