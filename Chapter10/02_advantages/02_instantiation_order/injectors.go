//+build wireinject

package main

import (
	"github.com/google/go-cloud/wire"
)

// The build tag makes sure the stub is not built in the final build.

func initializeDeps() *GetPersonHandler {
	wire.Build(wireSet)
	return nil
}
