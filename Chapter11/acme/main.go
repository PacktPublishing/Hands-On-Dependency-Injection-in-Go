package main

import (
	"context"
	"os"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/modules/exchange"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/modules/get"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/modules/list"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/modules/register"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch11/acme/internal/rest"
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
var wireSetWithoutConfig = wire.NewSet(
	// *exchange.Converter
	exchange.NewConverter,

	// *get.Getter
	get.NewGetter,

	// *list.Lister
	list.NewLister,

	// *register.Registerer
	wire.Bind(new(register.Exchanger), &exchange.Converter{}),
	register.NewRegisterer,

	// *rest.Server
	wire.Bind(new(rest.GetModel), &get.Getter{}),
	wire.Bind(new(rest.ListModel), &list.Lister{}),
	wire.Bind(new(rest.RegisterModel), &register.Registerer{}),
	rest.New,
)

var wireSet = wire.NewSet(
	wireSetWithoutConfig,

	// *config.Config
	config.Load,

	// *exchange.Converter
	wire.Bind(new(exchange.Config), &config.Config{}),

	// *get.Getter
	wire.Bind(new(get.Config), &config.Config{}),

	// *list.Lister
	wire.Bind(new(list.Config), &config.Config{}),

	// *register.Registerer
	wire.Bind(new(register.Config), &config.Config{}),

	// *rest.Server
	wire.Bind(new(rest.Config), &config.Config{}),
)
