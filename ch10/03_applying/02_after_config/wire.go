//+build ignore
// Code above this line should be ignored as it's not part of the example

//+build wireinject

package main

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/config"
	"github.com/google/go-cloud/wire"
)

// The build tag makes sure the stub is not built in the final build.

func initializeConfig() (*config.Config, error) {
	wire.Build(config.Load)
	return nil, nil
}
