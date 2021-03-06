package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/helstern/kommol/internal/bootstrap/core"
	"github.com/helstern/kommol/internal/bootstrap/logging"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object/get"
	"github.com/sarulabs/di/v2"
)

type router struct {
	key string
}

func (o *router) Get(ctn di.Container) *mux.Router {
	return ctn.Get(o.key).(*mux.Router)
}

func (o *router) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.key,
		Build: func(ctn di.Container) (interface{}, error) {
			router := mux.NewRouter()
			get.Routes(router, core.GetObjectProxy().Get(ctn), logging.GetLoggerFactory().Get(ctn))
			return router, nil
		},
	})
}

var (
	routerBootstap *router
)

func Router() *router {
	if routerBootstap == nil {
		routerBootstap = &router{
			key: builder.TypeName((*mux.Router)(nil)),
		}
	}

	return routerBootstap
}
