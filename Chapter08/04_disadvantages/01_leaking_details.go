package disadvantages

import (
	"errors"
	"strings"
)

type PeopleFilterConfig interface {
	DSN() string
}

func PeopleFilter(cfg PeopleFilterConfig, filter string) ([]Person, error) {
	// load people
	loader := &PersonLoader{}
	people, err := loader.LoadAll(cfg)
	if err != nil {
		return nil, err
	}

	// filter people
	out := []Person{}
	for _, person := range people {
		if strings.Contains(person.Name, filter) {
			out = append(out, person)
		}
	}

	return out, nil
}

type PersonLoaderConfig interface {
	DSN() string
}

type PersonLoader struct{}

func (p *PersonLoader) LoadAll(cfg PersonLoaderConfig) ([]Person, error) {
	return nil, errors.New("not implemented")
}

// Some data
type Person struct {
	Name string
}
