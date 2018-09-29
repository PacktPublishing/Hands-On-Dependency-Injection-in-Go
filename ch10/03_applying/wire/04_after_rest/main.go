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
	"github.com/google/go-cloud/wire"
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// start REST server
	server, err := initializeServer()
	if err != nil {
		os.Exit(-1)
	}

	server.Listen(ctx.Done())
}

// List of wire enabled objects
var wireSet = wire.NewSet(
	// *config.Config
	config.Load,

	// *exchange.Converter
	wire.Bind(new(exchange.Config), &config.Config{}),
	exchange.NewConverter,

	// *get.Getter
	wire.Bind(new(get.Config), &config.Config{}),
	get.NewGetter,

	// *list.Lister
	wire.Bind(new(list.Config), &config.Config{}),
	list.NewLister,

	// *register.Registerer
	wire.Bind(new(register.Config), &config.Config{}),
	wire.Bind(new(register.Exchanger), &exchange.Converter{}),
	register.NewRegisterer,

	// *rest.Server
	wire.Bind(new(rest.Config), &config.Config{}),
	wire.Bind(new(rest.GetModel), &get.Getter{}),
	wire.Bind(new(rest.ListModel), &list.Lister{}),
	wire.Bind(new(rest.RegisterModel), &register.Registerer{}),
	rest.New,
)
