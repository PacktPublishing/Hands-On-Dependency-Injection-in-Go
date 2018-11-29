package ocp

import (
	"net/http"
)

// a HTTP health check handler in short form
func healthCheckShort(resp http.ResponseWriter, _ *http.Request) {
	resp.WriteHeader(http.StatusNoContent)
}

func healthCheckShortUsage() {
	http.Handle("/health", http.HandlerFunc(healthCheckShort))
}
