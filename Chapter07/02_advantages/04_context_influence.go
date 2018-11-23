package advantages

import (
	"io"
	"net"
	"os"
)

func WriteLog(writer io.Writer, message string) error {
	_, err := writer.Write([]byte(message))
	return err
}

func Usage() {
	// Write to console
	WriteLog(os.Stdout, "Hello World!")

	// Write to file
	file, _ := os.Create("my-log.log")
	WriteLog(file, "Hello World!")

	// Write to TCP connection
	tcpPipe, _ := net.Dial("tcp", "127.0.0.1:1234")
	WriteLog(tcpPipe, "Hello World!")
}
