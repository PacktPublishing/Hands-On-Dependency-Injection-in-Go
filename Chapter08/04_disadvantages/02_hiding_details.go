package disadvantages

import (
	"strings"
)

type Loader interface {
	LoadAll() ([]Person, error)
}

func PeopleFilterV2(loader Loader, filter string) ([]Person, error) {
	// load people
	people, err := loader.LoadAll()
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
