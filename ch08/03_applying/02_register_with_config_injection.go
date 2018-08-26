// +build do-not-build

package applying

import (
	"context"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/exchange"
)

// NewRegisterer creates and initializes a Registerer
func NewRegisterer(cfg Config) *Registerer {
	return &Registerer{
		cfg: cfg,
	}
}

// Config is the configuration for the Registerer
type Config interface {
	Logger() logging.Logger
	RegistrationBasePrice() float64
}

// Registerer validates the supplied person, calculates the price in the requested currency and saves the result.
// It will return an error when:
// -the person object does not include all the fields
// -the currency is invalid
// -the exchange rate cannot be loaded
// -the data layer throws an error.
type Registerer struct {
	cfg Config
}

// get price in the requested currency
func (r *Registerer) getPrice(ctx context.Context, currency string) (float64, error) {
	converter := &exchange.Converter{}
	price, err := converter.Do(ctx, r.cfg.RegistrationBasePrice(), currency)
	if err != nil {
		r.logger().Warn("failed to convert the price. err: %s", err)
		return defaultPersonID, err
	}

	return price, nil
}

func (r *Registerer) logger() logging.Logger {
	return r.cfg.Logger()
}
