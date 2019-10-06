package gcp

import (
	"cloud.google.com/go/storage"
	"github.com/helstern/kommol/internal/core/object"
)

type ObjectWriterFactory func(bucket string, object string) object.Writer

func NewObjectWriterFactory(client *storage.Client) ObjectWriterFactory {
	return func(bucket string, object string) object.Writer {
		o := client.Bucket(bucket).Object(object)
		return NewObjectWriter(o)
	}
}
