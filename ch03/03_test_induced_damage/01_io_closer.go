package test_damage

import (
	"io"
)

func WriteAndClose(destination io.WriteCloser, contents string) error {
	defer destination.Close()

	_, err := destination.Write([]byte(contents))
	if err != nil {
		return err
	}

	return nil
}
