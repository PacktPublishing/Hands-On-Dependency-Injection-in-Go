package disadvantages

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveConfig(t *testing.T) {
	// inputs
	filename := "my-config.json"
	cfg := &Config{
		Host: "localhost",
		Port: 1234,
	}

	// monkey patch the file writer
	defer func(original func(filename string, data []byte, perm os.FileMode) error) {
		// restore the original
		writeFile = original
	}(writeFile)

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		// output error
		return nil
	}

	// call the function
	err := SaveConfig(filename, cfg)

	// validate the result
	assert.NoError(t, err)
}
