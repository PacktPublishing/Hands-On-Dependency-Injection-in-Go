package isp

import (
	"context"
	"errors"
)

func Encrypt(ctx context.Context, data []byte) ([]byte, error) {
	// As this operation make take too long, we need to be able to kill it
	stop := ctx.Done()
	result := make(chan []byte, 1)

	go func() {
		defer close(result)

		// pull the encryption key from context
		keyRaw := ctx.Value("encryption-key")
		if keyRaw == nil {
			panic("encryption key not found in context")
		}
		key := keyRaw.([]byte)

		// perform encryption
		ciperText := performEncryption(key, data)

		// signal complete by sending the result
		result <- ciperText
	}()

	select {
	case ciperText := <-result:
		// happy path
		return ciperText, nil

	case <-stop:
		// cancelled
		return nil, errors.New("operation cancelled")
	}
}

func performEncryption(key []byte, data []byte) []byte {
	// TODO: implement
	return nil
}
