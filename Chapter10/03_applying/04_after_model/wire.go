//+build ignore
// Code above this line should be ignored as it's not part of the example

//+build wireinject

package main

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/get"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/list"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/register"
	"github.com/google/go-cloud/wire"
)

// The build tag makes sure the stub is not built in the final build.

func initializeConfig() (*config.Config, error) {
	wire.Build(wireSet)
	return nil, nil
}

func initializeGetter() (*get.Getter, error) {
	wire.Build(wireSet)
	return nil, nil
}

func initializeLister() (*list.Lister, error) {
	wire.Build(wireSet)
	return nil, nil
}

func initializeRegisterer() (*register.Registerer, error) {
	wire.Build(wireSet)
	return nil, nil
}
