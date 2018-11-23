package advantages

import (
	"io"
)

func NewGeneratorV2(template io.Reader) *Generator {
	return &Generator{
		template: template,
	}
}

func (g *Generator) getStorage() Storage {
	if g.storage == nil {
		g.storage = &DefaultStorage{}
	}
	return g.storage
}

func (g *Generator) getRenderer() Renderer {
	if g.renderer == nil {
		g.renderer = &DefaultRenderer{}
	}
	return g.renderer
}

// Default implementation of Storage
type DefaultStorage struct{}

// Load implements Storage
func (d *DefaultStorage) Load() []interface{} {
	return nil
}

// Default implementation of Storage
type DefaultRenderer struct{}

// Load implements Renderer
func (d *DefaultRenderer) Render(template io.Reader, params ...interface{}) []byte {
	return nil
}
