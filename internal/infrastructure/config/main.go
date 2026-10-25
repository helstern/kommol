package config

import (
	"flag"
	"github.com/apex/log"
)

type Config struct {
	BindAddress  string
	GCP          GcpConfig
	LogLevel     log.Level
	NotFoundPage string
}

var (
	bind     = flag.String("bind", "127.0.0.1:8080", "Bind address")
	logLevel = flag.String("log-level", "info", "The logging level, defaults to info")
	notFound = flag.String("not-found-page", "", "Page served with 404 for a missing object, when the bucket names none; empty for a plain 404")
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
		BindAddress:  *bind,
		GCP:          NewGcpCliConfig(*gcpCredentials),
		LogLevel:     logLevel,
		NotFoundPage: *notFound,
	}
}
