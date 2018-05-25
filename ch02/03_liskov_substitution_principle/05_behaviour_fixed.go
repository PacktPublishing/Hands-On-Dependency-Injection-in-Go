package lsp

type ImmutableCollection interface {
	Get(index int) interface{}
}

type MutableCollection interface {
	ImmutableCollection
	Add(item interface{})
}

type ReadOnlyCollectionV2 struct {
	items []interface{}
}

func (ro *ReadOnlyCollectionV2) Get(index int) interface{} {
	return ro.items[index]
}

type CollectionImplV2 struct {
	ReadOnlyCollectionV2
}

func (c *CollectionImplV2) Add(item interface{}) {
	c.items = append(c.items, item)
}
