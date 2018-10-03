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
var wireSet = wire.NewSet(ProvideFetcher, ProvideCache)

// Providers
func ProvideFetcher(cache *Cache) *Fetcher {
	return &Fetcher{
		cache: cache,
	}
}

func ProvideCache() *Cache {
	return &Cache{}
}

type Cache struct{}

func (c *Cache) Get(key string) (string, error) {
	return "", errors.New("not implemented yet")
}

func (c *Cache) Set(key string, value string) error {
	return errors.New("not implemented")
}

type Fetcher struct {
	cache *Cache
}

func (f *Fetcher) GoFetch() (string, error) {
	return "", errors.New("not implemented yet")
}
