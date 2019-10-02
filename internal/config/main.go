package config

import (
	"flag"
	gcp "github.com/helstern/kommol/internal/core/gcp/config"
)

type Config struct {
	BindAddress string
	GCP         gcp.Config
}

var (
	bind    = flag.String("bind", "127.0.0.1:8080", "Bind address")
	verbose = flag.Bool("verbose", false, "Show access log")
)

// gcp config
var (
	gcpCredentials = flag.String("gcp.credentials", "", "The path to the keyfile. If not present, client will use your default application credentials.")
)

func NewCliConfig() Config {

	if !flag.Parsed() {
		flag.Parse()
	}

	return Config{
		BindAddress: *bind,
		GCP:         gcp.NewCliConfig(*gcpCredentials),
	}
}
