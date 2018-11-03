package advantages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ConfigSaver struct {
	FileWriter func(filename string, data []byte, perm os.FileMode) error
}

func (c ConfigSaver) Save(filename string, cfg *Config) error {
	// convert to JSON
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	// save file
	err = c.FileWriter(filename, data, 0666)
	if err != nil {
		log.Printf("failed to save file '%s' with err: %s", filename, err)
		return err
	}

	return nil
}

// Usage
func ConfigSaverUsage() {
	cfg := &Config{
		// build the config
	}

	saver := &ConfigSaver{
		FileWriter: ioutil.WriteFile,
	}

	err := saver.Save("myfile.json", cfg)
	if err != nil {
		fmt.Printf("failed with err: %s", err)
	}
}
