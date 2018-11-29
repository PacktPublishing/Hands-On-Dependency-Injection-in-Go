package disadvantages

import (
	"io"
)

type Generator struct{}

func (g *Generator) Generate(storage Storage, template io.Reader, destination io.Writer, renderer Renderer, formatter Formatter, params ...interface{}) {

}

type Storage interface {
	Load() []interface{}
}

type Renderer interface {
	Render(template io.Reader, params ...interface{}) []byte
}

type Formatter interface {
	Format([]byte) []byte
}
