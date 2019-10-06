package gcp

import (
	"cloud.google.com/go/storage"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/core/object/app"
	"net/url"
	"strings"
)

type ObjectRetrieverAdapter struct {
	createObjectWriter ObjectWriterFactory
}

func (this *ObjectRetrieverAdapter) Retrieve(opath string) (object.Writer, error) {

	oUrl, err := url.Parse(opath)
	if err != nil {
		return nil, err
	}

	return this.createObjectWriter(oUrl.Host, strings.Trim(oUrl.Path, "/")), nil
}

func NewObjectRetriever(createObjectWriter ObjectWriterFactory) app.ObjectRetriever {
	return &ObjectRetrieverAdapter{createObjectWriter: createObjectWriter}
}

func NewObjectRetrieverDefault(client *storage.Client) app.ObjectRetriever {
	return NewObjectRetriever(
		NewObjectWriterFactory(client),
	)
}
