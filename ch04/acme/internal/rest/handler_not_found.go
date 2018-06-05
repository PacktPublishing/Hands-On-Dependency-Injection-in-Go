package rest

import (
	"net/http"
)

func notFoundHandler(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusNotFound)
	_, _ = response.Write([]byte(`Not found`))
}
