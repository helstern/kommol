package object

import "github.com/helstern/kommol/internal/core"

type Object struct {
	Provider core.CloudProviders
	Path     []string
}
