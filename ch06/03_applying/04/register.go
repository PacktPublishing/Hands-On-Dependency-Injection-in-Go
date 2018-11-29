package rest

import (
	"net/http"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/04/data"
)

// RegisterModel will validate and save a registration
type RegisterModel interface {
	Do(in *data.Person) (int, error)
}

// RegisterHandler is the HTTP handler for the "Register" endpoint
// In this simplified example we are assuming all possible errors are user errors and returning "bad request" HTTP 400.
// There are some programmer errors possible but hopefully these will be caught in testing.
type RegisterHandler struct {
	registerer RegisterModel
}

// ServeHTTP implements http.Handler
func (h *RegisterHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// implementation goes here
}

// register endpoint request format
type registerRequest struct {
	// FullName of the person
	FullName string `json:"fullName"`
	// Phone of the person
	Phone string `json:"phone"`
	// Currency the wish to register in
	Currency string `json:"currency"`
}
