package needless_indirection

import (
	"net/http"
)

//go:generate mockery -name=MyMux -case underscore -testonly -inpkg -note @generated
type MyMux interface {
	Handle(pattern string, handler http.Handler)
	Handler(req *http.Request) (handler http.Handler, pattern string)
	ServeHTTP(resp http.ResponseWriter, req *http.Request)
}

// build HTTP handler routing
func buildRouter(mux MyMux) {
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
