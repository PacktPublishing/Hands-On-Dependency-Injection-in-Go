package register

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/05/data"
)

// Stub implementation so that the example compiles
type Registerer struct{}

func (r *Registerer) Do(in *data.Person) (int, error) {
	return 0, nil
}
