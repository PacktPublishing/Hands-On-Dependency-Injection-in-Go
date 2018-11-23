package advantages

import (
	"encoding/json"
	"net/http"
)

func HandlerV2(response http.ResponseWriter, request *http.Request) {
	garfield := &Animal{
		Type: "Cat",
		Name: "Garfield",
	}

	// encode as JSON and output
	outputAnimal(response, garfield)
}

func outputAnimal(response http.ResponseWriter, animal *Animal) {
	encoder := json.NewEncoder(response)
	err := encoder.Encode(animal)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Happy Path
	response.WriteHeader(http.StatusOK)
}
