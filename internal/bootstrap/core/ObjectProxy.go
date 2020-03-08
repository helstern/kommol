package core

import (
	"context"
	"github.com/helstern/kommol/internal/bootstrap/gcp"
	"github.com/helstern/kommol/internal/core"
	"github.com/helstern/kommol/internal/core/object/app"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"

	"github.com/sarulabs/di/v2"
)

type ObjectProxy struct {
	key string
}

func (o *ObjectProxy) Get(ctn di.Container) app.ObjectProxy {
	return ctn.Get(o.key).(app.ObjectProxy)
}

func (o *ObjectProxy) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.key,
		Build: func(ctn di.Container) (interface{}, error) {

			var resolvers = make(map[core.CloudProviders]app.ObjectProvider)
			resolvers[core.GCP] = gcp.GetObjectProvider().Get(ctn)

			return app.NewObjectProxy(resolvers), nil
		},
	})
}

var (
	bootstrapObjectProxy *ObjectProxy
)

func GetObjectProxy() *ObjectProxy {
	if bootstrapObjectProxy == nil {
		bootstrapObjectProxy = &ObjectProxy{
			key: builder.TypeName((*app.ObjectProxy)(nil)),
		}
	}

	return bootstrapObjectProxy
}
