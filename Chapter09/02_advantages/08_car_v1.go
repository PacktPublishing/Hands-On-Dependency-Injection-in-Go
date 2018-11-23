package advantages

type CarV1 struct {
	engine Engine
}

func (c *CarV1) Drive() {
	c.engine.Start()
	defer c.engine.Stop()

	c.engine.Drive()
}

type Engine interface {
	Start()
	Drive()
	Stop()
}
