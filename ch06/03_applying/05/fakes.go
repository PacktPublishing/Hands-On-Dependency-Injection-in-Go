package reset

import (
	"net/http"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/05/data"
)

func notFoundHandler(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusNotFound)
	_, _ = response.Write([]byte(`Not found`))
}

// Fake/Stub implementations to make the compiler happy

type Server struct {
	address string

	handlerGet      http.Handler
	handlerList     http.Handler
	handlerNotFound http.HandlerFunc
	handlerRegister http.Handler
}

func NewGetHandler(_ GetModel) *GetHandler {
	return &GetHandler{}
}

type GetModel interface {
	Do(ID int) (*data.Person, error)
}

type GetHandler struct{}

func (g *GetHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {}

func NewListHandler(_ ListModel) *ListHandler {
	return &ListHandler{}
}

type ListModel interface {
	Do() ([]*data.Person, error)
}

type ListHandler struct {
}

func (l *ListHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {}

func NewRegisterHandler(_ RegisterModel) *RegisterHandler {
	return &RegisterHandler{}
}

type RegisterModel interface {
	Do(in *data.Person) (int, error)
}

type RegisterHandler struct {
}

func (r *RegisterHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {}
