package isp

type myDB interface {
	GetItem(ID int) (Item, error)
	PutItem(item Item) error
}

type CacheV2 struct {
	ddb myDB
}

func (c *CacheV2) Get(key string) interface{} {
	// code removed

	// load from DB
	_, _ = c.ddb.GetItem(42)

	// code removed
	return nil
}

func (c *CacheV2) Set(key string, value interface{}) {
	// code removed

	// save from DB
	_ = c.ddb.PutItem(Item{})

	// code removed
}
