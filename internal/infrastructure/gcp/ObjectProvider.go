package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/core/http"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/core/object/app"
	"io"
	"time"
)

type ObjectProvider struct {
	app.ObjectProvider
	client  *storage.Client
	buckets BucketCache
}

func (this *ObjectProvider) fetchBucketWebsite(ctx context.Context, bucket string) (*storage.BucketWebsite, error) {
	websiteCfg, _ := this.buckets.Get(bucket)
	if websiteCfg == nil {
		cfgCtx, _ := context.WithTimeout(ctx, 3*time.Second)

		attrs, err := this.client.Bucket(bucket).Attrs(cfgCtx)
		if err != nil {
			return nil, err
		}

		if attrs.Website == nil {
			websiteCfg = &storage.BucketWebsite{
				MainPageSuffix: "",
				NotFoundPage:   "",
			}
		} else {
			websiteCfg = attrs.Website
		}
		this.buckets.Put(bucket, websiteCfg)
	}
	return websiteCfg, nil
}

func (this *ObjectProvider) WebsiteContainer(ctx context.Context, objectPath []string) (object.WebsiteContainer, error) {
	obj, err := ParsePath(objectPath)
	if err != nil {
		return EmptyWebsiteContainer(), err
	}

	websiteCfg, err := this.fetchBucketWebsite(ctx, obj.Bucket)
	if err != nil {
		return EmptyWebsiteContainer(), err
	}

	return NewWebsiteContainer(obj.Bucket, websiteCfg.MainPageSuffix)
}

func (this *ObjectProvider) Headers(ctx context.Context, objectPath []string) ([]http.Header, error) {
	obj, err := ParsePath(objectPath)
	if err != nil {
		return []http.Header{}, err
	}

	nextCtx, _ := context.WithTimeout(ctx, 3*time.Second)
	attrs, err := this.client.Bucket(obj.Bucket).Object(obj.Key).Attrs(nextCtx)
	if err != nil {
		return []http.Header{}, err
	}

	var headers []http.Header
	if attrs.ContentType != "" {
		headers = append(headers, http.NewContentTypeHeader(attrs.ContentType))
	}
	if attrs.ContentLanguage != "" {
		headers = append(headers, http.NewContentLanguageHeader(attrs.ContentLanguage))
	}
	if attrs.ContentEncoding != "" {
		headers = append(headers, http.NewContentEncodingHeader(attrs.ContentEncoding))
	}
	if attrs.ContentDisposition != "" {
		headers = append(headers, http.NewContentDispositionHeader(attrs.ContentDisposition))
	}
	if attrs.Size != 0 {
		headers = append(headers, http.NewContentLengthHeader(attrs.Size))
	}
	if attrs.CacheControl != "" {
		headers = append(headers, http.NewCacheControlHeader(attrs.CacheControl))
	}
	return headers, nil
}

func (this *ObjectProvider) Data(ctx context.Context, objectPath []string) (io.Reader, error) {
	obj, err := ParsePath(objectPath)
	if err != nil {
		return nil, err
	}

	reader, err := this.client.Bucket(obj.Bucket).Object(obj.Key).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func NewObjectProvider(client *storage.Client) ObjectProvider {
	return ObjectProvider{
		client:  client,
		buckets: NewBucketCache(),
	}
}
