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
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// load config
	cfg, err := config.Load(config.DefaultEnvVar)
	if err != nil {
		os.Exit(-1)
	}

	// build the exchanger
	exchanger := exchange.NewConverter(cfg)

	// build model layer
	getModel := get.NewGetter(cfg)
	listModel := list.NewLister(cfg)
	registerModel := register.NewRegisterer(cfg, exchanger)

	// start REST server
	server := rest.New(cfg, getModel, listModel, registerModel)
	server.Listen(ctx.Done())
}
