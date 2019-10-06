package core

import (
	"context"
	"github.com/helstern/kommol/internal/core/gcp"
	"github.com/helstern/kommol/internal/core/object/app"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/sarulabs/di/v2"
)

type objectProvider struct {
	key string
}

func (o *objectProvider) Get(ctn di.Container) app.ObjectRetriever {
	return ctn.Get(o.key).(app.ObjectRetriever)
}

func (o *objectProvider) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.key,
		Build: func(ctn di.Container) (interface{}, error) {
			client := StorageClient().Get(ctn)
			return gcp.NewObjectRetrieve(client), nil
		},
	})
}

var (
	objectProviderBootstrap *objectProvider
)

func ObjectRetriever() *objectProvider {
	if objectProviderBootstrap == nil {
		objectProviderBootstrap = &objectProvider{
			key: builder.TypeName((*gcp.ObjectRetriever)(nil)),
		}
	}

	return objectProviderBootstrap
}
