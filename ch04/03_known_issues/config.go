package known_issues

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/03_known_issues/currency"
)

type Config struct {
	DefaultCurrency currency.Currency `json:"default_currency"`
}
