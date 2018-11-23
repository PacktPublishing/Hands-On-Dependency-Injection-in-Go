package advantages

import (
	"io"
)

func NewGenerator(storage Storage, renderer Renderer, template io.Reader) *Generator {
	return &Generator{
		storage:  storage,
		renderer: renderer,
		template: template,
	}
}

type Generator struct {
	storage  Storage
	renderer Renderer
	template io.Reader
}

func (g *Generator) Generate(destination io.Writer, params ...interface{}) {

}

type Storage interface {
	Load() []interface{}
}

type Renderer interface {
	Render(template io.Reader, params ...interface{}) []byte
}
