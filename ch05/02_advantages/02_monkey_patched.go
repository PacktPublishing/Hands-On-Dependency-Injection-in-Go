package advantages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func SaveConfigPatched(filename string, cfg *Config) error {
	// convert to JSON
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	// save file
	err = writeFile(filename, data, 0666)
	if err != nil {
		log.Printf("failed to save file '%s' with err: %s", filename, err)
		return err
	}

	return nil
}

// Custom type that allows us to Monkey Patch
var writeFile = ioutil.WriteFile

// Usage
func SaveConfigPatchedUsage() {
	cfg := &Config{
		// build the config
	}

	err := SaveConfigPatched("myfile.json", cfg)
	if err != nil {
		fmt.Printf("failed with err: %s", err)
	}
}
