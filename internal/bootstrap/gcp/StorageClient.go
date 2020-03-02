package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/bootstrap/config"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/sarulabs/di/v2"
	"google.golang.org/api/option"
)

type StorageClient struct {
	key string
}

func (o *StorageClient) Get(ctn di.Container) *storage.Client {
	return ctn.Get(o.key).(*storage.Client)
}

func (o *StorageClient) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.key,
		Build: func(ctn di.Container) (interface{}, error) {
			appConfig := config.Get(ctn)
			credentials := appConfig.GCP.GetCredentialFilePath()

			if credentials != "" {
				return storage.NewClient(ctx, option.WithCredentialsFile(credentials), option.WithScopes("https://www.googleapis.com/auth/devstorage.read_only"))
			}

			return storage.NewClient(ctx, option.WithScopes("https://www.googleapis.com/auth/devstorage.read_only"))
		},
	})
}

var (
	bootstrapStorageClient *StorageClient
)

func GetStorageClient() *StorageClient {
	if bootstrapStorageClient == nil {
		bootstrapStorageClient = &StorageClient{
			key: builder.TypeName((*storage.Client)(nil)),
		}
	}

	return bootstrapStorageClient
}
