package config

import (
	"flag"
	"github.com/apex/log"
)

type Config struct {
	BindAddress string
	GCP         GcpConfig
	LogLevel    log.Level
}

var (
	bind     = flag.String("bind", "127.0.0.1:8080", "Bind address")
	logLevel = flag.String("log-level", "info", "The logging level, defaults to info")
)

// gcp config
var (
	gcpCredentials = flag.String("gcp.credentials", "", "The path to the keyfile. If not present, client will use your default application credentials.")
)

func NewCliConfig() Config {

	if !flag.Parsed() {
		flag.Parse()
	}

	logLevel, _ := log.ParseLevel(*logLevel)

	return Config{
		BindAddress: *bind,
		GCP:         NewGcpCliConfig(*gcpCredentials),
		LogLevel:    logLevel,
	}
}
