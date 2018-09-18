// +build bad

package payment

import (
	"errors"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch01/02_code_smells/04_tight_coupling/circular_dependenciescular_dependencies/config"
)

// Currency is custom type for currency
type Currency string

// Processor processes payments
type Processor struct {
	Config *config.Config
}

// Pay makes a payment in the default currency
func (p *Processor) Pay(amount float64) error {
	// TODO: implement me
	return errors.New("not implemented yet")
}
