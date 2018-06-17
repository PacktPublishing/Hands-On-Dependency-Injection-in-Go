package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New will create and initialize the server
func New(address string) *Server {
	return &Server{
		address:         address,
		handlerGet:      &GetHandler{},
		handlerList:     &ListHandler{},
		handlerNotFound: notFoundHandler,
		handlerRegister: &RegisterHandler{},
	}
}

// Server is the HTTP REST server
type Server struct {
	address string
	server  *http.Server

	handlerGet      http.Handler
	handlerList     http.Handler
	handlerNotFound http.HandlerFunc
	handlerRegister http.Handler
}

// Listen will start a HTTP rest for this service
func (s *Server) Listen(stop <-chan struct{}) {
	router := s.buildRouter()

	// create the HTTP server
	s.server = &http.Server{
		Handler: router,
		Addr:    s.address,
	}

	// listen for shutdown
	go func() {
		// wait for shutdown signal
		<-stop

		_ = s.server.Close()
	}()

	// start the HTTP server
	_ = s.server.ListenAndServe()
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
