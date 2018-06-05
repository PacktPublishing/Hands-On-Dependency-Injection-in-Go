package test_damage

import (
	"encoding/json"
	"io"
)

func PrintAsJSON(destination io.Writer, plant Plant) error {
	bytes, err := json.Marshal(plant)
	if err != nil {
		return err
	}

	destination.Write(bytes)
	return nil
}

type Plant struct {
	Name string
}
