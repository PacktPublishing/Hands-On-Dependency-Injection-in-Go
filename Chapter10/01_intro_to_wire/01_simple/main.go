package main

import (
	"errors"
	"fmt"

	"github.com/google/go-cloud/wire"
)

func main() {
	f := initializeDeps()

	result, err := f.GoFetch()
	fmt.Printf("Result: %s / %s", result, err)
}

// list of wire enabled dependencies
var wireSet = wire.NewSet(ProvideFetcher)

// Provider
func ProvideFetcher() *Fetcher {
	return &Fetcher{}
}

// Object being "provided"
type Fetcher struct {
}

func (f *Fetcher) GoFetch() (string, error) {
	return "", errors.New("not implemented yet")
}
