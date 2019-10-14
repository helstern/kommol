package gcp

import (
	"context"
	"github.com/helstern/kommol/internal/bootstrap/gcp/api"
	"github.com/helstern/kommol/internal/bootstrap/gcp/core"
	"github.com/sarulabs/di/v2"
)

func Module(ctx context.Context, builder *di.Builder) error {
	var err error

	err = api.Object().Module(ctx, builder)
	if err != nil {
		return err
	}

	err = core.ObjectRetriever().Module(ctx, builder)
	if err != nil {
		return err
	}

	err = core.StorageClient().Module(ctx, builder)
	if err != nil {
		return err
	}

	return nil
}
