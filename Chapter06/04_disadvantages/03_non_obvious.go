package disadvantages

import (
	"errors"
)

// NewClient creates and initialises the client
func NewClient(service DepService) Client {
	return &clientImpl{
		service: service,
	}
}

// Client is the exported API
type Client interface {
	DoSomethingUseful() (bool, error)
}

// implement Client
type clientImpl struct {
	service DepService
}

func (c *clientImpl) DoSomethingUseful() (bool, error) {
	// this function does something useful
	return false, errors.New("not implemented")
}

type DepService interface {
	DoSomethingElse()
}
