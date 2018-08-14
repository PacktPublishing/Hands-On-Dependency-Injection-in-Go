package advantages

import (
	"encoding/json"
	"net/http"
)

func HandlerV1(response http.ResponseWriter, request *http.Request) {
	garfield := &Animal{
		Type: "Cat",
		Name: "Garfield",
	}

	// encode as JSON and output
	encoder := json.NewEncoder(response)
	err := encoder.Encode(garfield)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}

type Animal struct {
	Type string
	Name string
}
