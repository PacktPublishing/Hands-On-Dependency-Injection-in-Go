package isp

import (
	"context"
)

func UseEncryptV2() {
	// create a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// call the function
	_, _ = EncryptV2(ctx, ctx, []byte("my data"))
}
