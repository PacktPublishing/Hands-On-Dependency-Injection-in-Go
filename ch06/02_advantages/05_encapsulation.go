package advantages

import (
	"errors"
)

func (c *CarV2) FillPetrolTank() error {
	// use the engine
	if c.engine.IsRunning() {
		return errors.New("cannot fill the tank while the engine is running")
	}

	// fill the tank!
	return c.fill()
}

func (c CarV2) fill() error {
	// TODO: implement
	return nil
}
