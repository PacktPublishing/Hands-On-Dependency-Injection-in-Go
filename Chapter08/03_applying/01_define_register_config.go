// +build do-not-build

package applying

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/logging"
)

// Config is the configuration for the Registerer
type Config interface {
	Logger() *logging.LoggerStdOut
	BasePrice() float64
}
