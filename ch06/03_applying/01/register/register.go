package register

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/01/data"
)

// Registerer validates the supplied person, calculates the price in the requested currency and saves the result.
// It will return an error when:
// -the person object does not include all the fields
// -the currency is invalid
// -the exchange rate cannot be loaded
// -the data layer throws an error.
type Registerer struct {
}

// Do is API for this struct
func (r *Registerer) Do(in *data.Person) (int, error) {
	// fake implementation
	return 0, nil
}
