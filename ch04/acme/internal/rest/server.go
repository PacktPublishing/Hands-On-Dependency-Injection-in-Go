package rest

import (
	"net/http"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/common/logging"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch04/acme/internal/config"
	"github.com/gorilla/mux"
)

// New will create and initialize the server
func New() *Server {
	return &Server{
		handlerGet:      &GetHandler{},
		handlerList:     &ListHandler{},
		handlerNotFound: notFoundHandler,
		handlerRegister: &RegisterHandler{},
	}
}

// Server is the HTTP REST server
type Server struct {
	handlerGet      http.Handler
	handlerList     http.Handler
	handlerNotFound http.HandlerFunc
	handlerRegister http.Handler
}

// Listen will start a HTTP rest for this service
func (s *Server) Listen() {
	router := s.buildRouter()

	// create and start a HTTP server
	server := &http.Server{
		Handler: router,
		Addr:    config.App.Address,
	}
	logging.Warn("", server.ListenAndServe())
}

// configure the endpoints to handlers
func (s *Server) buildRouter() http.Handler {
	router := mux.NewRouter()

	// map URL endpoints to HTTP handlers
	router.Handle("/person/{id}/", s.handlerGet).Methods("GET")
	router.Handle("/person/list", s.handlerList).Methods("GET")
	router.Handle("/person/register", s.handlerRegister).Methods("POST")

	// convert a "catch all" not found handler
	router.NotFoundHandler = s.handlerNotFound

	return router
}
