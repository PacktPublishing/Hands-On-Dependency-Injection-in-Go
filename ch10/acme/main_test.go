package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	// start a context with a max execution time
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// start test server
	serverAddress := startTestServer(t, ctx)

	// build and send request
	payload := bytes.NewBufferString(`
{
	"fullName": "Bob",
	"phone": "0123456789",
	"currency": "AUD"
}
`)

	req, err := http.NewRequest("POST", serverAddress+"/person/register", payload)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	// validate expectations
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.NotEmpty(t, resp.Header.Get("Location"))
}

func TestGet(t *testing.T) {
	// start a context with a max execution time
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// start test server
	serverAddress := startTestServer(t, ctx)

	// build and send request
	req, err := http.NewRequest("GET", serverAddress+"/person/1/", nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	// validate expectations
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestList(t *testing.T) {
	// start a context with a max execution time
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// start test server
	serverAddress := startTestServer(t, ctx)

	// build and send request
	req, err := http.NewRequest("GET", serverAddress+"/person/list", nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	// validate expectations
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func startTestServer(t *testing.T, ctx context.Context) string {
	// load the standard config (from the ENV)
	cfg, err := config.Load()
	require.NoError(t, err)

	// get a free port (so tests can run concurrently)
	port, err := getFreePort()
	require.NoError(t, err)

	// override config port with free one
	cfg.Address = net.JoinHostPort("0.0.0.0", port)

	// start the test server on a random port
	go func() {
		// start REST server
		server := initializeServerCustomConfig(cfg, cfg, cfg, cfg, cfg)
		server.Listen(ctx.Done())
	}()

	// give the server a chance to start
	<-time.After(100 * time.Millisecond)

	// return the address of the test server
	return "http://" + cfg.Address
}

func getFreePort() (string, error) {
	for attempt := 0; attempt <= 10; attempt++ {
		addr := net.JoinHostPort("", "0")
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			continue
		}

		port, err := getPort(listener.Addr())
		if err != nil {
			continue
		}

		// close/free the port
		tcpListener := listener.(*net.TCPListener)
		cErr := tcpListener.Close()
		if cErr == nil {
			file, fErr := tcpListener.File()
			if fErr == nil {
				// ignore any errors cleaning up the file
				_ = file.Close()
			}
			return port, nil
		}
	}

	return "", errors.New("no free ports")
}

func getPort(addr fmt.Stringer) (string, error) {
	actualAddress := addr.String()
	_, port, err := net.SplitHostPort(actualAddress)
	return port, err
}
