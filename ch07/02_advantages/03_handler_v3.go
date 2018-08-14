package advantages

import (
	"encoding/json"
	"net/http"
)

func HandlerV3(response http.ResponseWriter, request *http.Request) {
	garfield := &Animal{
		Type: "Cat",
		Name: "Garfield",
	}

	// encode as JSON and output
	outputJSON(response, garfield)
}

func outputJSON(response http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(response)
	err := encoder.Encode(data)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Happy Path
	response.WriteHeader(http.StatusOK)
}
