package logging

import (
	"context"
	"github.com/helstern/kommol/internal/core/logging/app"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/helstern/kommol/internal/infrastructure/logging"
	"github.com/sarulabs/di/v2"
)

type LoggerFactory struct {
	key string
}

func (o *LoggerFactory) Get(ctn di.Container) app.LoggerFactory {
	return ctn.Get(o.key).(app.LoggerFactory)
}

func (o *LoggerFactory) Module(ctx context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: o.key,
		Build: func(ctn di.Container) (interface{}, error) {
			return app.LoggerFactory(logging.NewLoggerWithFields), nil
		},
	})
}

var (
	bootstrapLoggerFactory *LoggerFactory
)

func GetLoggerFactory() *LoggerFactory {
	if bootstrapLoggerFactory == nil {
		bootstrapLoggerFactory = &LoggerFactory{
			key: builder.TypeName((app.LoggerFactory)(nil)),
		}
	}

	return bootstrapLoggerFactory
}
