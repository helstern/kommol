package main

import (
	"context"
	apexLog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/helstern/kommol/internal/bootstrap"
	httpBootstrap "github.com/helstern/kommol/internal/bootstrap/http"
	"github.com/sarulabs/di/v2"
	"log"
	"net/http"
	"os"
)

func main() {

	ctx := context.Background()
	server := getServer(ctx)

	//ctxx := apexLog.WithFields(apexLog.Fields{
	//	"file": "something.png",
	//	"type": "image/png",
	//	"user": "tobi",
	//})
	//ctxx.Info("upload")
	//ctxx.Info("upload complete")
	//ctxx.Warn("upload retry")
	//ctxx.Errorf("failed to upload %s", "img.png")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getServer(ctx context.Context) *http.Server {
	apexLog.SetHandler(text.New(os.Stdout))
	builder, _ := di.NewBuilder()

	_ = bootstrap.Setup(ctx, builder)
	ctn := builder.Build()
	return httpBootstrap.Server().Get(ctn)
}
