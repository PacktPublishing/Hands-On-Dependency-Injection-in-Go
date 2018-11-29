//+build ignore
// Code above this line should be ignored as it's not part of the example

package main

import (
	"context"
	"os"
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
