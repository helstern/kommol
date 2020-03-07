package gcp

import (
	"context"
	"github.com/helstern/kommol/internal/bootstrap/logging"
	"github.com/helstern/kommol/internal/core/object/app"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	"github.com/sarulabs/di/v2"
)

type ObjectProvider struct {
	key string
}

func (o *ObjectProvider) Get(ctn di.Container) app.ObjectProvider {
	return ctn.Get(o.key).(app.ObjectProvider)
}

func (o *ObjectProvider) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.key,
		Build: func(ctn di.Container) (interface{}, error) {
			client := GetStorageClient().Get(ctn)
			logger := logging.GetLoggerFactory().Get(ctn)
			return gcp.NewObjectProvider(client, logger), nil
		},
	})
}

var (
	bootstrapObjectProvider *ObjectProvider
)

func GetObjectProvider() *ObjectProvider {
	if bootstrapObjectProvider == nil {
		bootstrapObjectProvider = &ObjectProvider{
			key: builder.TypeName((*gcp.ObjectProvider)(nil)),
		}
	}

	return bootstrapObjectProvider
}
