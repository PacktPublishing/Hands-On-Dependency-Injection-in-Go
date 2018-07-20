package reset

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/05/get"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/05/list"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/03_applying/05/register"
)

// New will create and initialize the server
func New(address string) *Server {
	return &Server{
		address:         address,
		handlerGet:      NewGetHandler(&get.Getter{}),
		handlerList:     NewListHandler(&list.Lister{}),
		handlerNotFound: notFoundHandler,
		handlerRegister: NewRegisterHandler(&register.Registerer{}),
	}
}
