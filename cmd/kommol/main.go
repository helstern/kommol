package main

import (
	"context"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/helstern/kommol/internal/bootstrap"
	configBootstrap "github.com/helstern/kommol/internal/bootstrap/config"
	httpBootstrap "github.com/helstern/kommol/internal/bootstrap/http"
	loggingBootstrap "github.com/helstern/kommol/internal/bootstrap/logging"
	logging "github.com/helstern/kommol/internal/core/logging/app"
	"github.com/helstern/kommol/internal/infrastructure/config"
	"github.com/sarulabs/di/v2"

	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	container := buildDI(ctx)

	appConfig := configBootstrap.Get(container)
	preRun(appConfig)

	server := httpBootstrap.Server().Get(container)
	loggerFactory := loggingBootstrap.GetLoggerFactory().Get(container)
	run(server, logging.ContextLogger(ctx, loggerFactory))
}

func preRun(config config.Config) {
	log.SetHandler(text.New(os.Stdout))
	log.SetLevel(config.LogLevel)
}

func run(server *http.Server, logger logging.Logger) {
	logger.Info("starting server")
	if err := server.ListenAndServe(); err != nil {
		logger.WithError(err).Fatal("server failure")
	}
}

func buildDI(ctx context.Context) di.Container {
	builder, _ := di.NewBuilder()

	_ = bootstrap.Setup(ctx, builder)
	ctn := builder.Build()
	return ctn
}
