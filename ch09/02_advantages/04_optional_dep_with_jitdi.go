package advantages

func NewLoaderWithJIT(ds Datastore) *LoaderWithJIT {
	return &LoaderWithJIT{
		datastore: ds,
	}
}

type LoaderWithJIT struct {
	// required private dependency
	datastore Datastore

	// optional cache
	OptionalCache Cache
}

func (l *LoaderWithJIT) Load(ID int) (*Animal, error) {
	// attempt to load from cache
	output := l.cache().Get(ID)
	if output != nil {
		// return cached value
		return output, nil
	}

	// load from data store
	output, err := l.datastore.Load(ID)
	if err != nil {
		return nil, err
	}

	// cache the loaded value
	l.cache().Put(ID, output)

	// output the result
	return output, nil
}

func (l *LoaderWithJIT) cache() Cache {
	if l.OptionalCache == nil {
		l.OptionalCache = &noopCache{}
	}

	return l.OptionalCache
}

// NO-OP implementation of the cache
type noopCache struct {
	// intentionally blank
}

func (n *noopCache) Get(ID int) *Animal {
	// intentionally does nothing
	return nil
}

func (n *noopCache) Put(ID int, value *Animal) {
	// intentionally does nothing
}
