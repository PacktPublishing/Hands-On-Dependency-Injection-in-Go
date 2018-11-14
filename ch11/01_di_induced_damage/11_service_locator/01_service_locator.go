package service_locator

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		deps: map[string]interface{}{},
	}
}

type ServiceLocator struct {
	deps map[string]interface{}
}

// Store or map a dependency to a key
func (s *ServiceLocator) Store(key string, dep interface{}) {
	s.deps[key] = dep
}

// Retrieve a dependency by key
func (s *ServiceLocator) Get(key string) interface{} {
	return s.deps[key]
}
