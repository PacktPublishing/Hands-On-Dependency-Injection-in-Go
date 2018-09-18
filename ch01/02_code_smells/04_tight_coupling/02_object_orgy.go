package _4_tight_coupling

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type PageLoader struct {
}

func (o *PageLoader) LoadPage(url string) ([]byte, error) {
	b := newFetcher()

	// check cache
	payload, err := b.cache.Get(url)
	if err == nil {
		// found in cache
		return payload, nil
	}

	// call upstream
	resp, err := b.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// extract data from HTTP response
	payload, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// save to cache asynchronously
	go func(key string, value []byte) {
		b.cache.Set(key, value)
	}(url, payload)

	// return
	return payload, nil
}

type Fetcher struct {
	httpClient http.Client
	cache      *Cache
}

func newFetcher() *Fetcher {
	return &Fetcher{}
}

type Cache struct {
	// not implemented
}

func (c *Cache) Get(key string) ([]byte, error) {
	// not implemented
	return nil, errors.New("not implemented")
}

func (c *Cache) Set(key string, data []byte) error {
	// not implemented
	return errors.New("not implemented")
}
