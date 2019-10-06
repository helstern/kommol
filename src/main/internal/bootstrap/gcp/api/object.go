package api

import (
	"context"
	"github.com/helstern/kommol/internal/bootstrap/gcp/core"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	objectPkg "github.com/helstern/kommol/internal/presentation/api/gcp/object"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object/operations"
	"github.com/sarulabs/di/v2"
)

type object struct {
	keyGetHandler string
}

func (o *object) Get(ctn di.Container) objectPkg.GetHandler {
	return ctn.Get(o.keyGetHandler).(objectPkg.GetHandler)
}

func (o *object) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.keyGetHandler,
		Build: func(ctn di.Container) (interface{}, error) {
			client := core.StorageClient().Get(ctn)
			objectRetriever := gcp.NewObjectRetrieverDefault(client)

			return operations.NewGetHandler(objectRetriever), nil
		},
	})
}

var (
	objectBootstrap *object
)

func Object() *object {
	if objectBootstrap == nil {
		objectBootstrap = &object{
			keyGetHandler: builder.TypeName((*objectPkg.GetHandler)(nil)),
		}
	}
	return objectBootstrap
}
