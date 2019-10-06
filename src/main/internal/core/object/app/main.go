package app

import "github.com/helstern/kommol/internal/core/object"

type ObjectRetriever interface {
	Retrieve(path string) (object.Writer, error)
}
