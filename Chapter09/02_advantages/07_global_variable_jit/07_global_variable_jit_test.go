package global_variable_jit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaver_Do(t *testing.T) {
	// input
	carol := &User{
		Name:     "Carol",
		Password: "IamKing",
	}

	// mocks/stubs
	stubStorage := &StubUserStorage{}

	// do call
	saver := &Saver{
		storage: stubStorage,
	}
	resultErr := saver.Do(carol)

	// validate
	assert.NotEqual(t, resultErr, "unexpected error")
}

// Stub implementation of UserStorage
type StubUserStorage struct{}

func (s *StubUserStorage) Save(_ *User) error {
	// return "happy path"
	return nil
}
