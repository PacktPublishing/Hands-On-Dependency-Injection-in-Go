package list

import (
	"errors"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/acme/internal/modules/data"
)

var (
	// error thrown when there are no people in the database
	errPeopleNotFound = errors.New("no people found")
)

// Lister will attempt to load all people in the database.
// It can return an error caused by the data layer
type Lister struct {
}

// Do will load the people from the data layer
func (l *Lister) Do() ([]*data.Person, error) {
	// load all people
	people, err := l.load()
	if err != nil {
		return nil, err
	}

	if len(people) == 0 {
		// special processing for 0 people returned
		return nil, errPeopleNotFound
	}

	return people, nil
}

// load all people
func (l *Lister) load() ([]*data.Person, error) {
	people, err := loader()
	if err != nil {
		if err == data.ErrNotFound {
			// By converting the error we are encapsulating the implementation details from our users.
			return nil, errPeopleNotFound
		}
		return nil, err
	}

	return people, nil
}

// this function as a variable allows us to Monkey Patch during testing
var loader = data.LoadAll
