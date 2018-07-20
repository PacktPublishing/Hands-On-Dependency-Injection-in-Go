package list

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/05/data"
)

// Stub implementation so that the example compiles
type Lister struct{}

func (l *Lister) Do() ([]*data.Person, error) {
	return nil, nil
}
