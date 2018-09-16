package jit_injection

import (
	"errors"
	"net/http"
)

func NewLoadPersonHandler(logic LoadPersonLogic) *LoadPersonHandler {
	return &LoadPersonHandler{
		businessLogic: logic,
	}
}

type LoadPersonHandler struct {
	businessLogic LoadPersonLogic
}

func (h *LoadPersonHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	requestedID, err := h.extractInputFromRequest(request)

	output, err := h.businessLogic.Load(requestedID)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.writeOutput(response, output)
}

// extract the person ID from the request
func (h *LoadPersonHandler) extractInputFromRequest(request *http.Request) (int, error) {
	return 0, errors.New("not implemented yet")
}

// convert person to JSON and write to the HTTP response
func (h *LoadPersonHandler) writeOutput(writer http.ResponseWriter, person Person) {
	// not implemented yet
}

type LoadPersonLogic interface {
	// Load person by supplied ID
	Load(ID int) (Person, error)
}
