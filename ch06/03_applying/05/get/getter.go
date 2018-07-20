package get

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/05/data"
)

// Stub implementation so that the example compiles
type Getter struct{}

func (g *Getter) Do(ID int) (*data.Person, error) {
	return nil, nil
}
