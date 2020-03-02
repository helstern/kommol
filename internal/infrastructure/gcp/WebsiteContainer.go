package gcp

import "github.com/helstern/kommol/internal/core/object"
import "github.com/helstern/kommol/internal/core"

func EmptyWebsiteContainer() object.WebsiteContainer {
	return object.WebsiteContainer{Provider: core.GCP}
}

func NewWebsiteContainer(bucket string, indexFile string) (object.WebsiteContainer, error) {
	return object.NewWebsiteContainer(core.GCP, []string{bucket}, indexFile)
}
