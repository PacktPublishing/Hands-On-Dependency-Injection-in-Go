package config_coupling

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/03_known_issues/02_config_coupling/currency"
)

type Config struct {
	DefaultCurrency currency.Currency `json:"default_currency"`
}
