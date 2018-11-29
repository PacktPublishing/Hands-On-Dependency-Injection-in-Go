package disadvantages

import (
	"database/sql"
	"errors"
)

// Load people from the database
type PersonLoader struct {
}

func (d *PersonLoader) Load(db *sql.DB, ID int) (*Person, error) {
	return nil, errors.New("not implemented")
}

func (d *PersonLoader) LoadAll(db *sql.DB) ([]*Person, error) {
	return nil, errors.New("not implemented")
}

type Person struct {
	Name string
	Age  int
}
