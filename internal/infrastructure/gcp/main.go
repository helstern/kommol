package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/core/object"
	"io"
)

type Object struct {
	Bucket string
	Key    string
}

func (o *Object) GetPath() string {
	return "gs://" + o.Bucket + "/" + o.Key
}

func (o *Object) Attrs(ctx context.Context, client *storage.Client) (attrs *storage.ObjectAttrs, err error) {
	return client.Bucket(o.Bucket).Object(o.Key).Attrs(ctx)
}

func (o *Object) NewReader(ctx context.Context, client *storage.Client) (io.Reader, error) {
	return client.Bucket(o.Bucket).Object(o.Key).NewReader(ctx)
}

type GcpObjectRetriever interface {
	RetrieveObject(object *Object) (object.Writer, error)
}

type GoogleClientObjectRetriever struct {
	GcpObjectRetriever
	Client *storage.Client
}

func (this *GoogleClientObjectRetriever) RetrieveObject(object *Object) (object.Writer, error) {
	return NewObjectWriterAdapter(this.Client, object), nil
}
