package register

import (
	"context"
	"errors"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/data"
)

const (
	// default person id (returned on error)
	defaultPersonID = 0
)

var (
	// validation errors
	errNameMissing     = errors.New("name is missing")
	errPhoneMissing    = errors.New("name is missing")
	errCurrencyMissing = errors.New("currency is missing")
	errInvalidCurrency = errors.New("currency is invalid, supported types are AUD, CNY, EUR, GBP, JPY, MYR, SGD, USD")

	// a little trick to make checking for supported currencies easier
	supportedCurrencies = map[string]struct{}{
		"AUD": {},
		"CNY": {},
		"EUR": {},
		"GBP": {},
		"JPY": {},
		"MYR": {},
		"SGD": {},
		"USD": {},
	}
)

// NewRegisterer creates and initializes a Registerer
func NewRegisterer(cfg Config, exchanger Exchanger) *Registerer {
	return &Registerer{
		cfg:       cfg,
		exchanger: exchanger,
	}
}

// Exchanger will convert from one currency to another
type Exchanger interface {
	// Exchange will perform the conversion
	Exchange(ctx context.Context, basePrice float64, currency string) (float64, error)
}

// Config is the configuration for the Registerer
type Config interface {
	Logger() logging.Logger
	RegistrationBasePrice() float64
	DataDSN() string
}

// Registerer validates the supplied person, calculates the price in the requested currency and saves the result.
// It will return an error when:
// -the person object does not include all the fields
// -the currency is invalid
// -the exchange rate cannot be loaded
// -the data layer throws an error.
type Registerer struct {
	cfg       Config
	exchanger Exchanger
}

// Do is API for this struct
func (r *Registerer) Do(ctx context.Context, in *data.Person) (int, error) {
	// validate the request
	err := r.validateInput(in)
	if err != nil {
		r.logger().Warn("input validation failed with err: %s", err)
		return defaultPersonID, err
	}

	// get price in the requested currency
	price, err := r.getPrice(ctx, in.Currency)
	if err != nil {
		return defaultPersonID, err
	}

	// save registration
	id, err := r.save(ctx, in, price)
	if err != nil {
		// no need to log here as we expect the data layer to do so
		return defaultPersonID, err
	}

	return id, nil
}

// validate input and return error on fail
func (r *Registerer) validateInput(in *data.Person) error {
	if in.FullName == "" {
		return errNameMissing
	}
	if in.Phone == "" {
		return errPhoneMissing
	}
	if in.Currency == "" {
		return errCurrencyMissing
	}

	if _, found := supportedCurrencies[in.Currency]; !found {
		return errInvalidCurrency
	}

	// happy path
	return nil
}

// get price in the requested currency
func (r *Registerer) getPrice(ctx context.Context, currency string) (float64, error) {
	price, err := r.exchanger.Exchange(ctx, r.cfg.RegistrationBasePrice(), currency)
	if err != nil {
		r.logger().Warn("failed to convert the price. err: %s", err)
		return defaultPersonID, err
	}

	return price, nil
}

// save the registration
func (r *Registerer) save(ctx context.Context, in *data.Person, price float64) (int, error) {
	person := &data.Person{
		FullName: in.FullName,
		Phone:    in.Phone,
		Currency: in.Currency,
		Price:    price,
	}
	return saver(ctx, r.cfg, person)
}

func (r *Registerer) logger() logging.Logger {
	return r.cfg.Logger()
}

// this function as a variable allows us to Monkey Patch during testing
var saver = data.Save
