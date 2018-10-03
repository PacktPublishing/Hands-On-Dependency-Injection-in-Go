//+build ignore
// Code above this line should be ignored as it's not part of the example

package main

import (
	"context"
	"os"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/exchange"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/get"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/list"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/register"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/rest"
	"go.uber.org/dig"
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// build DIG container
	container := BuildContainer()

	// start REST server
	err := container.Invoke(func(server *rest.Server) {
		server.Listen(ctx.Done())
	})

	if err != nil {
		os.Exit(-1)
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.Load)
	container.Provide(exchange.NewConverter)
	container.Provide(get.NewGetter)
	container.Provide(list.NewLister)
	container.Provide(register.NewRegisterer)
	container.Provide(rest.New)

	return container
}
