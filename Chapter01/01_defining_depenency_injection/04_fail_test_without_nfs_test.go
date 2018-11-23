package defining_depenency_injection

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSavePerson_nfsAlwaysFails(t *testing.T) {
	// input
	in := &Person{
		Name:  "Sophia",
		Phone: "0123456789",
	}

	// mock the NFS
	mockNFS := &mockSaver{}
	mockNFS.On("Save", mock.Anything).Return(errors.New("save failed")).Once()

	// Call Save
	resultErr := SavePerson(in, mockNFS)

	// validate result
	assert.Error(t, resultErr)
	assert.True(t, mockNFS.AssertExpectations(t))
}
