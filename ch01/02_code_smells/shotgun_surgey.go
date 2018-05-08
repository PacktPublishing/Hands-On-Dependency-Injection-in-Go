package code_smells

import (
	"io"
	"database/sql"
)

// Person data object
type Person struct {
	Name string
	Phone string
}

// Renderer will render a person to the supplied writer
type Renderer struct{}

func (r Renderer) render(p *Person, output io.Writer) {
	// output the person
}

// Validator will validate the supplied person has all the required fields
type Validator struct{}

func (v Validator) validate(p *Person) error {
	// validate the person
	return nil
}

// Saver will save the supplied person to the DB
type Saver struct{}

func (s *Saver) Save(db *sql.DB, p *Person) {
	// save the person to db
}
