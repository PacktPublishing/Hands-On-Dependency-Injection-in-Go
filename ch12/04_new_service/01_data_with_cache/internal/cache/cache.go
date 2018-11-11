package cache

import (
	"errors"
)

type Cache struct{}

func (c *Cache) Get(key string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (c *Cache) Set(key string, data []byte) error {
	return errors.New("not implemented")
}
