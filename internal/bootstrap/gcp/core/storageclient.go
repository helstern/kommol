package core

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/bootstrap/config"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/sarulabs/di/v2"
	"google.golang.org/api/option"
)

type storageClient struct {
	key string
}

func (o *storageClient) Get(ctn di.Container) *storage.Client {
	return ctn.Get(o.key).(*storage.Client)
}

func (o *storageClient) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.key,
		Build: func(ctn di.Container) (interface{}, error) {
			appConfig := config.Get(ctn)
			credentials := appConfig.GCP.GetCredentialFilePath()

			if credentials != "" {
				return storage.NewClient(ctx, option.WithCredentialsFile(credentials))
			}

			return storage.NewClient(ctx)
		},
	})
}

var (
	storageClientBootstrap *storageClient
)

func StorageClient() *storageClient {
	if storageClientBootstrap == nil {
		storageClientBootstrap = &storageClient{
			key: builder.TypeName((*storage.Client)(nil)),
		}
	}

	return storageClientBootstrap
}
