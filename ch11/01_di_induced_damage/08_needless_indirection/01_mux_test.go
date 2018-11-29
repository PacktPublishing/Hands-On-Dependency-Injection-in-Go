package needless_indirection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildRouter(t *testing.T) {
	// build mock
	mockRouter := &MockMyMux{}
	mockRouter.On("Handle", "/get", &getEndpoint{}).Once()
	mockRouter.On("Handle", "/list", &listEndpoint{}).Once()
	mockRouter.On("Handle", "/save", &saveEndpoint{}).Once()

	// call function
	buildRouter(mockRouter)

	// assert expectations
	assert.True(t, mockRouter.AssertExpectations(t))
}
