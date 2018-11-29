package advantages

import (
	"errors"
)

func NewLoader(ds Datastore, cache Cache) *MyLoader {
	return &MyLoader{
		ds:    ds,
		cache: cache,
	}
}

type MyLoader struct {
	ds    Datastore
	cache Cache
}

func (l *MyLoader) LoadAll() ([]interface{}, error) {
	return nil, errors.New("not implemented")
}
