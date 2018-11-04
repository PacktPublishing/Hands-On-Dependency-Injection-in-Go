package advantages

import (
	"errors"
)

func (c *CarV2) FillPetrolTankV2(engine Engine) error {
	// use the engine
	if engine.IsRunning() {
		return errors.New("cannot fill the tank while the engine is running")
	}

	// fill the tank!
	return c.fill()
}
