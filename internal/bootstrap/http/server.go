package http

import (
	"context"
	"github.com/helstern/kommol/internal/bootstrap/config"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/sarulabs/di/v2"
	"net/http"
)

type server struct {
	key string
}

func (o *server) Get(ctn di.Container) *http.Server {
	return ctn.Get(o.key).(*http.Server)
}

func (o *server) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.key,
		Build: func(ctn di.Container) (interface{}, error) {
			router := Router().Get(ctn)
			cfg := config.Get(ctn)

			return &http.Server{
				Addr:    cfg.BindAddress,
				Handler: router,
			}, nil
		},
	})
}

var (
	serverBootstrap *server
)

func Server() *server {
	if serverBootstrap == nil {
		serverBootstrap = &server{
			key: builder.TypeName((*http.Server)(nil)),
		}
	}

	return serverBootstrap
}
