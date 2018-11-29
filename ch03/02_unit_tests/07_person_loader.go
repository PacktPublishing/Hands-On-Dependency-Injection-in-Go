package unit_tests

import (
	"errors"
)

var ErrNotFound = errors.New("person not found")

type Person struct {
	Name string
}

//go:generate mockery -name PersonLoader -testonly -inpkg -case=underscore
type PersonLoader interface {
	Load(ID int) (*Person, error)
}

func LoadPersonName(loader PersonLoader, ID int) (string, error) {
	person, err := loader.Load(ID)
	if err != nil {
		return "", err
	}

	return person.Name, nil
}
