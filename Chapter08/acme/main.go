package main

import (
	"context"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/exchange"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/get"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/list"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/modules/register"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch08/acme/internal/rest"
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// build the exchanger
	exchanger := exchange.NewConverter(config.App)

	// build model layer
	getModel := get.NewGetter(config.App)
	listModel := list.NewLister(config.App)
	registerModel := register.NewRegisterer(config.App, exchanger)

	// start REST server
	server := rest.New(config.App, getModel, listModel, registerModel)
	server.Listen(ctx.Done())
}
