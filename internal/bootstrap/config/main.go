package config

import (
	"context"
	"github.com/helstern/kommol/internal/config"
	"github.com/helstern/kommol/internal/infrastructure/di/builder"
	"github.com/sarulabs/di/v2"
)

var key = builder.TypeName((*config.Config)(nil))

func Get(ctn di.Container) config.Config {
	return ctn.Get(key).(config.Config)
}

func Module(_ context.Context, builder *di.Builder) error {
	return builder.Add(di.Def{
		Name: key,
		Build: func(ctn di.Container) (interface{}, error) {
			return config.NewCliConfig(), nil
		},
	})
}
