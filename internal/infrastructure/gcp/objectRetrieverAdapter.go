package gcp

import (
	"cloud.google.com/go/storage"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/core/object/app"
	"net/url"
	"strings"
)

type ObjectRetrieverAdapter struct {
	GcpObjectRetriever
}

func (this *ObjectRetrieverAdapter) Retrieve(opath string) (object.Writer, error) {

	oUrl, err := url.Parse(opath)
	if err != nil {
		return nil, err
	}

	bucket := oUrl.Host
	key := strings.Trim(oUrl.Path, "/")

	gcpObject := &Object{
		Bucket: bucket,
		Key:    key,
	}

	return this.RetrieveObject(gcpObject)
}

func NewObjectRetrieverAdapter(retriever GcpObjectRetriever) *ObjectRetrieverAdapter {
	return &ObjectRetrieverAdapter{
		GcpObjectRetriever: retriever,
	}
}

func NewObjectRetriever(client *storage.Client) app.ObjectRetriever {
	return NewObjectRetrieverAdapter(&GoogleClientObjectRetriever{
		Client:  client,
		buckets: BucketCache{},
	})
}
