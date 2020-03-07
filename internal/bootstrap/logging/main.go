package logging

import (
	"context"
	"github.com/sarulabs/di/v2"
)

func Module(ctx context.Context, builder *di.Builder) error {
	var err error

	err = GetLoggerFactory().Module(ctx, builder)
	if err != nil {
		return err
	}

	return nil
}
