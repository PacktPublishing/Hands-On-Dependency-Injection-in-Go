package advantages

import (
	"errors"
)

type Engine interface {
	Start()
	IncreasePower()
	DecreasePower()
	Stop()
	IsRunning() bool
}

type Car struct {
	Engine Engine
}

func (c *Car) Drive() error {
	if c.Engine == nil {
		return errors.New("engine ie missing")
	}

	// use the engine
	c.Engine.Start()
	c.Engine.IncreasePower()

	return nil
}

func (c *Car) Stop() error {
	if c.Engine == nil {
		return errors.New("engine ie missing")
	}

	// use the engine
	c.Engine.DecreasePower()
	c.Engine.Stop()

	return nil
}
