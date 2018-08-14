package method_injection

import (
	"fmt"
	"io"
	"os"
	"time"
)

// TimeStampWriterV3 will output the supplied message to writer preceded with a timestamp
func TimestampWriterV3(writer io.Writer, message string) {
	if writer == nil {
		// default to Standard Out
		writer = os.Stdout
	}

	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("%s -> %s", timestamp, message)
}
