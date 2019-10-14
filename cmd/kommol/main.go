package main

import (
	"context"
	"github.com/helstern/kommol/internal/bootstrap"
	httpBootstrap "github.com/helstern/kommol/internal/bootstrap/http"
	"github.com/sarulabs/di/v2"
	"log"
	"net/http"
)

func main() {

	ctx := context.Background()
	server := getServer(ctx)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getServer(ctx context.Context) *http.Server {

	builder, _ := di.NewBuilder()

	_ = bootstrap.Setup(ctx, builder)
	ctn := builder.Build()
	return httpBootstrap.Server().Get(ctn)
}
