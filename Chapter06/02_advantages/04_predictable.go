package advantages

import (
	"errors"
)

func NewCarV2(engine Engine) (*CarV2, error) {
	if engine == nil {
		return nil, errors.New("invalid engine supplied")
	}

	return &CarV2{
		engine: engine,
	}, nil
}

type CarV2 struct {
	engine Engine
}

func (c *CarV2) Drive() error {
	// use the engine
	c.engine.Start()
	c.engine.IncreasePower()

	return nil
}

func (c *CarV2) Stop() error {
	// use the engine
	c.engine.DecreasePower()
	c.engine.Stop()

	return nil
}
