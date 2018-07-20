package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/01/data"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/01/register"
)

// RegisterHandler is the HTTP handler for the "Register" endpoint
// In this simplified example we are assuming all possible errors are user errors and returning "bad request" HTTP 400.
// There are some programmer errors possible but hopefully these will be caught in testing.
type RegisterHandler struct {
}

// ServeHTTP implements http.Handler
func (h *RegisterHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// extract payload from request
	requestPayload, err := h.extractPayload(request)
	if err != nil {
		// output error
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// register person
	id, err := h.register(requestPayload)
	if err != nil {
		// not need to log here as we can expect other layers to do so
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	// happy path
	response.Header().Add("Location", fmt.Sprintf("/person/%d/", id))
	response.WriteHeader(http.StatusCreated)
}

// extract payload from request
func (h *RegisterHandler) extractPayload(request *http.Request) (*registerRequest, error) {
	requestPayload := &registerRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(requestPayload)
	if err != nil {
		return nil, err
	}

	return requestPayload, nil
}

// call the logic layer
func (h *RegisterHandler) register(requestPayload *registerRequest) (int, error) {
	person := &data.Person{
		FullName: requestPayload.FullName,
		Phone:    requestPayload.Phone,
		Currency: requestPayload.Currency,
	}

	registerer := &register.Registerer{}
	return registerer.Do(person)
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
