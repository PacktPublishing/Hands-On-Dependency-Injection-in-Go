package method_injection

import (
	"errors"
	"fmt"
	"io"
	"time"
)

// TimeStampWriterV2 will output the supplied message to writer preceded with a timestamp
func TimestampWriterV2(writer io.Writer, message string) error {
	if writer == nil {
		return errors.New("writer cannot be nil")
	}

	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("%s -> %s", timestamp, message)

	return nil
}
