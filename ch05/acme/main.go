package main

import (
	"context"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch05/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch05/acme/internal/rest"
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// start REST server
	server := rest.New(config.App.Address)
	server.Listen(ctx.Done())
}
