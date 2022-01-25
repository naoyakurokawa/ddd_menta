package middleware

import (
	"context"

	"github.com/naoyakurokawa/ddd_menta/config"
)

func getAppErrorFromCtx(ctx context.Context) error {
	val := ctx.Value(config.AppErrorKey)
	if val == nil {
		return nil
	}

	err, ok := val.(error)
	if !ok {
		return nil
	}

	return err
}
