package advantages

func NewLoaderWithoutJIT(ds Datastore) *LoaderWithoutJIT {
	return &LoaderWithoutJIT{
		datastore: ds,
	}
}

type LoaderWithoutJIT struct {
	// required private dependency
	datastore Datastore

	// optional cache
	OptionalCache Cache
}

func (l *LoaderWithoutJIT) Load(ID int) (*Animal, error) {
	var output *Animal
	var err error

	// attempt to load from cache
	if l.OptionalCache != nil {
		output = l.OptionalCache.Get(ID)
		if output != nil {
			// return cached value
			return output, nil
		}
	}

	// load from data store
	output, err = l.datastore.Load(ID)
	if err != nil {
		return nil, err
	}

	// cache the loaded value
	if l.OptionalCache != nil {
		l.OptionalCache.Put(ID, output)
	}

	// output the result
	return output, nil
}

type Cache interface {
	Get(ID int) *Animal
	Put(ID int, value *Animal)
}

type Datastore interface {
	Load(ID int) (*Animal, error)
	Save(ID int, value *Animal) error
}

type Animal struct {
	// some data fields go here
}
