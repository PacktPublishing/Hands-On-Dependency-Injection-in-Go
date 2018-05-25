package ocp

import (
	"io"
	"net/http"
)

func BuildOutputOCPFail(response http.ResponseWriter, format string, person Person) {
	var err error

	switch format {
	case "csv":
		err = outputCSV(response, person)

	case "json":
		err = outputJSON(response, person)
	}

	if err != nil {
		// output a server error and quit
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
}

// output the person as CSV and return error when failing to do so
func outputCSV(writer io.Writer, person Person) error {
	// TODO: implement
	return nil
}

// output the person as JSON and return error when failing to do so
func outputJSON(writer io.Writer, person Person) error {
	// TODO: implement
	return nil
}

// A data transfer object that represents a person
type Person struct {
	Name  string
	Email string
}
