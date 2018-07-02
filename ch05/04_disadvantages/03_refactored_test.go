package disadvantages

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveConfig_refactored(t *testing.T) {
	// inputs
	filename := "my-config.json"
	cfg := &Config{
		Host: "localhost",
		Port: 1234,
	}

	// monkey patch the file writer
	defer restoreWriteFile(writeFile)

	writeFile = mockWriteFile(nil)

	// call the function
	err := SaveConfig(filename, cfg)

	// validate the result
	assert.NoError(t, err)
}

func mockWriteFile(result error) func(filename string, data []byte, perm os.FileMode) error {
	return func(filename string, data []byte, perm os.FileMode) error {
		return result
	}
}

// remove the restore function to reduce from 3 lines to 1
func restoreWriteFile(original func(filename string, data []byte, perm os.FileMode) error) {
	// restore the original
	writeFile = original
}
