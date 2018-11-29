package disadvantages

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func SaveConfig(filename string, cfg *Config) error {
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

// Custom type that allows
var writeFile = ioutil.WriteFile

type Config struct {
	Host string
	Port int
}
