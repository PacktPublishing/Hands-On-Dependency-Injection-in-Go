package rest

import (
	"context"
	"net"
)

func getOpenPort() (string, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}

	address := listener.Addr().String()
	listener.Close()

	return address, nil
}

func startServer(ctx context.Context) (string, error) {
	// get open port
	address, err := getOpenPort()
	if err != nil {
		return "", err
	}

	// start a server
	server := New(address)
	go server.Listen(ctx.Done())

	// wait for server to be ready
	dialer := &net.Dialer{}
	for {
		conn, _ := dialer.DialContext(ctx, "tcp", address)
		if conn != nil {
			defer conn.Close()

			return address, nil
		}

		select {
		case <-ctx.Done():
			return "", ctx.Err()

		default:
			// try again
		}
	}

	return address, nil
}
