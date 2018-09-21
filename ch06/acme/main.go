package main

import (
	"context"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/acme/internal/modules/get"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/acme/internal/modules/list"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/acme/internal/modules/register"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/acme/internal/rest"
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// build model layer
	getModel := &get.Getter{}
	listModel := &list.Lister{}
	registerModel := &register.Registerer{}

	// start REST server
	server := rest.New(config.App.Address, getModel, listModel, registerModel)
	server.Listen(ctx.Done())
}
