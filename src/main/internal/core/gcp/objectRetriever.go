package gcp

import (
	"cloud.google.com/go/storage"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/core/object/app"
	"path"
	"strings"
)

type ObjectRetriever struct {
	client *storage.Client
}

func (this *ObjectRetriever) Retrieve(opath string) (object.Writer, error) {
	cleanPath := path.Clean(opath)
	segments := strings.Split(cleanPath, "/")
	bucket, key := segments[0], strings.Join(segments[1:], "/")

	o := (*this.client).Bucket(bucket).Object(key)
	return NewObjectWriter(o), nil
}

func NewObjectRetrieve(client *storage.Client) app.ObjectRetriever {
	return &ObjectRetriever{client: client}
}
