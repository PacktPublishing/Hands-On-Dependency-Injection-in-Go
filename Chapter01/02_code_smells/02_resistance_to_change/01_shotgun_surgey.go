package _2_resistance_to_change

import (
	"database/sql"
	"io"
)

// Renderer will render a person to the supplied writer
type Renderer struct{}

func (r Renderer) render(name, phone string, output io.Writer) {
	// output the person
}

// Validator will validate the supplied person has all the required fields
type Validator struct{}

func (v Validator) validate(name, phone string) error {
	// validate the person
	return nil
}

// Saver will save the supplied person to the DB
type Saver struct{}

func (s *Saver) Save(db *sql.DB, name, phone string) {
	// save the person to db
}
