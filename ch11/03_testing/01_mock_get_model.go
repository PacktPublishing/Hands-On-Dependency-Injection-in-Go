package testing

import (
	"github.com/stretchr/testify/mock"
)

type MockGetModel struct {
	mock.Mock
}

func (_m *MockGetModel) Do(ID int) (*Person, error) {
	outputs := _m.Called(ID)

	if outputs.Get(0) != nil {
		return outputs.Get(0).(*Person), outputs.Error(1)
	}

	return nil, outputs.Error(1)
}

type Person struct {
	ID       int
	FullName string
	Phone    string
	Currency string
	Price    float64
}
