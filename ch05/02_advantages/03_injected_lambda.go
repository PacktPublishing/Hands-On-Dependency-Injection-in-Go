package advantages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func SaveConfigInjected(writer fileWriter, filename string, cfg *Config) error {
	// convert to JSON
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	// save file
	err = writer(filename, data, 0666)
	if err != nil {
		log.Printf("failed to save file '%s' with err: %s", filename, err)
		return err
	}

	return nil
}

// This custom type is not strictly needed but it does make the function
// signature a little cleaner
type fileWriter func(filename string, data []byte, perm os.FileMode) error

// Usage
func SaveConfigInjectedUsage() {
	cfg := &Config{
		// build the config
	}

	err := SaveConfigInjected(ioutil.WriteFile, "myfile.json", cfg)
	if err != nil {
		fmt.Printf("failed with err: %s", err)
	}
}
