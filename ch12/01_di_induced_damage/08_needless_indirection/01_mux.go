package needless_indirection

import (
	"net/http"
)

// build HTTP handler routing
func buildRouter(mux *http.ServeMux) {
	mux.Handle("/get", &getEndpoint{})
	mux.Handle("/list", &listEndpoint{})
	mux.Handle("/save", &saveEndpoint{})
}

type getEndpoint struct{}

func (*getEndpoint) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	// not implemented
}

type listEndpoint struct{}

func (*listEndpoint) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	// not implemented
}

type saveEndpoint struct{}

func (*saveEndpoint) ServeHTTP(_ http.ResponseWriter, _ *http.Request) {
	// not implemented
}
