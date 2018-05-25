package isp

import (
	"context"
)

type Item struct {
	Key     string
	Payload []byte
}

type FatDbInterface interface {
	BatchGetItem(IDs ...int) ([]Item, error)
	BatchGetItemWithContext(ctx context.Context, IDs ...int) ([]Item, error)

	BatchPutItem(items ...Item) error
	BatchPutItemWithContext(ctx context.Context, items ...Item) error

	DeleteItem(ID int) error
	DeleteItemWithContext(ctx context.Context, item Item) error

	GetItem(ID int) (Item, error)
	GetItemWithContext(ctx context.Context, ID int) (Item, error)

	PutItem(item Item) error
	PutItemWithContext(ctx context.Context, item Item) error

	Query(query string, args ...interface{}) ([]Item, error)
	QueryWithContext(ctx context.Context, query string, args ...interface{}) ([]Item, error)

	UpdateItem(item Item) error
	UpdateItemWithContext(ctx context.Context, item Item) error
}

type Cache struct {
	db FatDbInterface
}

func (c *Cache) Get(key string) interface{} {
	// code removed

	// load from DB
	_, _ = c.db.GetItem(42)

	// code removed
	return nil
}

func (c *Cache) Set(key string, value interface{}) {
	// code removed

	// save to DB
	_ = c.db.PutItem(Item{})

	// code removed
}
