package rest

import (
	"fmt"
	"net/http"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/common/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/modules/register"
	"github.com/gorilla/schema"
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
	response.WriteHeader(http.StatusCreated)
	response.Header().Add("Location", fmt.Sprintf("/person/%d/", id))
}

// extract payload from request
func (h *RegisterHandler) extractPayload(request *http.Request) (*registerRequest, error) {
	// parse the HTTP request body
	err := request.ParseForm()
	if err != nil {
		// load and return error
		logging.Warn("[register] bad request. err: %s", err)
		return nil, err
	}

	// build a struct to hold the request data
	requestPayload := &registerRequest{}

	// decode from request.Form.Values into a struct
	decoder := schema.NewDecoder()
	err = decoder.Decode(requestPayload, request.PostForm)
	if err != nil {
		// log and return error
		logging.Error("[register] failed to decode request. err: %s", err)
		return nil, err
	}

	return requestPayload, nil
}

// call the logic layer
func (h *RegisterHandler) register(requestPayload *registerRequest) (int, error) {
	person := &register.Person{
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
