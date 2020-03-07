package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/core"
	"github.com/helstern/kommol/internal/core/http"
	logging "github.com/helstern/kommol/internal/core/logging/app"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/core/object/app"
	"io"
	"time"
)

type ObjectProvider struct {
	app.ObjectProvider
	client  *storage.Client
	buckets BucketCache
	logging logging.LoggerFactory
}

func (this ObjectProvider) fetchBucketWebsite(ctx context.Context, bucket string) (*storage.BucketWebsite, error) {
	providerName, _ := core.GCP.Name()
	logger := logging.ContextLogger(ctx, this.logging).WithFields(logging.Fields{
		"provider": providerName,
		"bucket":   bucket,
	})

	websiteCfg, _ := this.buckets.Get(bucket)
	if websiteCfg != nil {
		logger.Debug("read bucket configuration from cache")
		return websiteCfg, nil
	}

	logger.Debug("requesting bucket configuration")
	attrs, err := this.client.Bucket(bucket).Attrs(ctx)
	if err != nil {
		return nil, err
	}

	if attrs.Website == nil {
		logger.Debug("website configuration not found")
		websiteCfg = &storage.BucketWebsite{
			MainPageSuffix: "",
			NotFoundPage:   "",
		}
	} else {
		websiteCfg = attrs.Website
	}
	this.buckets.Put(bucket, websiteCfg)

	return websiteCfg, nil
}

func (this ObjectProvider) WebsiteContainer(ctx context.Context, objectPath []string) (object.WebsiteContainer, error) {
	obj, err := ParsePath(objectPath)
	if err != nil {
		return EmptyWebsiteContainer(), err
	}

	providerName, _ := core.GCP.Name()
	logging.ContextLogger(ctx, this.logging).WithFields(logging.Fields{
		"providerName": providerName,
		"bucket":       obj.Bucket,
	}).Debug("reading bucket configuration")

	fetchCtx, _ := context.WithTimeout(ctx, 3*time.Second)
	websiteCfg, err := this.fetchBucketWebsite(fetchCtx, obj.Bucket)
	if err != nil {
		return EmptyWebsiteContainer(), err
	}

	return NewWebsiteContainer(obj.Bucket, websiteCfg.MainPageSuffix)
}

func (this ObjectProvider) Headers(ctx context.Context, objectPath []string) ([]http.Header, error) {
	obj, err := ParsePath(objectPath)
	if err != nil {
		return []http.Header{}, err
	}

	providerName, _ := core.GCP.Name()
	logger := logging.ContextLogger(ctx, this.logging).WithFields(logging.Fields{
		"provider": providerName,
		"bucket":   obj.Bucket,
		"key":      obj.Key,
	})
	nextCtx, _ := context.WithTimeout(ctx, 3*time.Second)

	logger.Debug("reading object attributes")
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

func (this ObjectProvider) Data(ctx context.Context, objectPath []string) (io.Reader, error) {
	obj, err := ParsePath(objectPath)
	if err != nil {
		return nil, err
	}

	providerName, _ := core.GCP.Name()
	logging.ContextLogger(ctx, this.logging).WithFields(logging.Fields{
		"provider": providerName,
		"bucket":   obj.Bucket,
		"key":      obj.Key,
	}).Debug("reading object data")
	reader, err := this.client.Bucket(obj.Bucket).Object(obj.Key).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func NewObjectProvider(client *storage.Client, logging logging.LoggerFactory) app.ObjectProvider {
	return ObjectProvider{
		client:  client,
		buckets: NewBucketCache(),
		logging: logging,
	}
}
