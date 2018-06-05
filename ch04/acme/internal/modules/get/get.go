package get

import (
	"errors"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/modules/data"
)

var (
	// error thrown when the requested person is not in the database
	errPersonNotFound = errors.New("person not found")
)

// Person is the data transfer object (DTO) for this package
type Person struct {
	ID       int
	FullName string
	Phone    string
	Currency string
	Price    float64
}

// Getter will attempt to load a person.
// It can return an error caused by the data layer or when the requested person is not found
type Getter struct {
}

// Do will perform the get
func (g *Getter) Do(in int) (*Person, error) {
	// load person from the data layer
	person, err := g.load(in)
	if err != nil {
		return nil, err
	}

	// build output
	return g.buildOutput(person), nil
}

// load person from the data layer
func (g *Getter) load(in int) (*data.Person, error) {
	person, err := data.Load(in)
	if err != nil {
		if err == data.ErrNotFound {
			// By converting the error we are encapsulating the implementation details from our users.
			return nil, errPersonNotFound
		}
		return nil, err
	}

	return person, err
}

// Convert from the data layer DTO to the DTO for this layer.
// Thus encapsulating the implementation details from our users.
func (g *Getter) buildOutput(person *data.Person) *Person {
	return &Person{
		ID:       person.ID,
		FullName: person.FullName,
		Phone:    person.Phone,
		Currency: person.Currency,
		Price:    person.Price,
	}
}
