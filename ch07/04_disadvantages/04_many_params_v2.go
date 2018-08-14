package disadvantages

import (
	"io"
)

func NewGeneratorV2(storage Storage, renderer Renderer, formatter Formatter) *GeneratorV2 {
	return &GeneratorV2{
		storage:   storage,
		renderer:  renderer,
		formatter: formatter,
	}
}

type GeneratorV2 struct {
	storage   Storage
	renderer  Renderer
	formatter Formatter
}

func (g *GeneratorV2) Generate(template io.Reader, destination io.Writer, params ...interface{}) {

}
