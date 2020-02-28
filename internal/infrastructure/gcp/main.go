package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/core/object"
	"io"
	"time"
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
	Client  *storage.Client
	buckets BucketCache
}

func (this *GoogleClientObjectRetriever) resolveIndexObject(bucket string) *Object {
	var websiteCfg *storage.BucketWebsite

	websiteCfg, _ = this.buckets.Get(bucket)
	if websiteCfg == nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		attrs, _ := this.Client.Bucket(bucket).Attrs(ctx)
		if attrs.Website == nil {
			websiteCfg = &storage.BucketWebsite{
				MainPageSuffix: "/",
				NotFoundPage:   "",
			}
		} else {
			websiteCfg = attrs.Website
		}
		this.buckets.Put(bucket, websiteCfg)
	}

	return &Object{
		Bucket: bucket,
		Key:    websiteCfg.MainPageSuffix,
	}
}

func (this *GoogleClientObjectRetriever) resolveObject(object *Object) *Object {
	if object.Key == "" || object.Key == "/" {
		return this.resolveIndexObject(object.Bucket)
	}

	return object
}

func (this *GoogleClientObjectRetriever) RetrieveObject(object *Object) (object.Writer, error) {
	resolvedObject := this.resolveObject(object)
	return NewObjectWriterAdapter(this.Client, resolvedObject), nil
}
