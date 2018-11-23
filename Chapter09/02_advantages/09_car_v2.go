package advantages

type CarV2 struct {
	engine Engine
}

func (c *CarV2) Drive() {
	engine := c.getEngine()

	engine.Start()
	defer engine.Stop()

	engine.Drive()
}

func (c *CarV2) getEngine() Engine {
	if c.engine == nil {
		c.engine = newEngine()
	}

	return c.engine
}

func newEngine() Engine {
	// not implemented
	return nil
}
