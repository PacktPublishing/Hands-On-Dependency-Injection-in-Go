// +build do-not-build

package applying

import (
	"context"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/data"
)

// GetterConfig is the configuration for Getter
type GetterConfig interface {
	Logger() logging.Logger
}

// Getter will attempt to load a person.
// It can return an error caused by the data layer or when the requested person is not found
type Getter struct {
	cfg GetterConfig
}

// Do will perform the get
func (g *Getter) Do(ID int) (*data.Person, error) {
	// load person from the data layer
	person, err := loader(context.TODO(), ID)
	if err != nil {
		if err == data.ErrNotFound {
			// By converting the error we are hiding the implementation details from our users.
			return nil, errPersonNotFound
		}
		return nil, err
	}

	return person, err
}

// this function as a variable allows us to Monkey Patch during testing
var loader = data.Load
