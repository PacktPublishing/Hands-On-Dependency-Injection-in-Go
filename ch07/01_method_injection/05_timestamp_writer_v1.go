package method_injection

import (
	"fmt"
	"io"
	"time"
)

// TimeStampWriterV1 will output the supplied message to writer preceded with a timestamp
func TimestampWriterV1(writer io.Writer, message string) {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("%s -> %s", timestamp, message)
}
