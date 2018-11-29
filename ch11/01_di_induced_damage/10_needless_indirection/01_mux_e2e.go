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

func (*getEndpoint) ServeHTTP(resp http.ResponseWriter, _ *http.Request) {
	_, _ = resp.Write([]byte(`Hi from Get!`))
}

type listEndpoint struct{}

func (*listEndpoint) ServeHTTP(resp http.ResponseWriter, _ *http.Request) {
	_, _ = resp.Write([]byte(`Hi from List!`))
}

type saveEndpoint struct{}

func (*saveEndpoint) ServeHTTP(resp http.ResponseWriter, _ *http.Request) {
	_, _ = resp.Write([]byte(`Hi from Save!`))
}
