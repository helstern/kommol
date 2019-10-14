package bootstrap

import (
	"context"
	"github.com/helstern/kommol/internal/bootstrap/config"
	"github.com/helstern/kommol/internal/bootstrap/gcp"
	"github.com/helstern/kommol/internal/bootstrap/http"
	"github.com/sarulabs/di/v2"
)

func Setup(ctx context.Context, builder *di.Builder) error {
	var err error

	err = http.Module(ctx, builder)
	if err != nil {
		return err
	}

	err = config.Module(ctx, builder)
	if err != nil {
		return err
	}

	err = gcp.Module(ctx, builder)
	if err != nil {
		return err
	}

	return nil
}
