package rest

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/02/register"
)

// RegisterHandler is the HTTP handler for the "Register" endpoint
// In this simplified example we are assuming all possible errors are user errors and returning "bad request" HTTP 400.
// There are some programmer errors possible but hopefully these will be caught in testing.
type RegisterHandler struct {
	registerer *register.Registerer
}
