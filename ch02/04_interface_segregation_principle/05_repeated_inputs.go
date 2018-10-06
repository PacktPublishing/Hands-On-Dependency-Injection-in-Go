package isp

import (
	"context"
)

func UseEncryptV2() {
	// create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// store the key
	ctx = context.WithValue(ctx, "encryption-key", "-secret-")

	// call the function
	_, _ = EncryptV2(ctx, ctx, []byte("my data"))
}
