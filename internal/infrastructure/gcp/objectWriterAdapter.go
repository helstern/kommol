package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/infrastructure/http/headers"
	"io"
	"net/http"
)

func ModifyHeaders(attrs *storage.ObjectAttrs, headerMap http.Header) {
	if attrs.ContentType != "" {
		headerMap.Set(headers.ContentType, attrs.ContentType)
	}

	if attrs.ContentLanguage != "" {
		headerMap.Set(headers.ContentLanguage, attrs.ContentLanguage)
	}

	if attrs.ContentEncoding != "" {
		headerMap.Set(headers.ContentEncoding, attrs.ContentEncoding)
	}

	if attrs.ContentDisposition != "" {
		headerMap.Set(headers.ContentDisposition, attrs.ContentDisposition)
	}

	if attrs.Size != 0 {
		headerMap.Set(headers.ContentLength, headers.FormatInt(attrs.Size))
	}

	if attrs.CacheControl != "" {
		headerMap.Set(headers.CacheControl, attrs.CacheControl)
	}
}

type ObjectWriterAdapter struct {
	object.Writer
	client *storage.Client
	object *Object
}

func (this *ObjectWriterAdapter) ModifyHeaders(ctx context.Context, headerMap http.Header) error {
	attr, err := this.object.Attrs(ctx, this.client)
	if err != nil {
		return err
	}

	ModifyHeaders(attr, headerMap)
	return nil
}

func (this *ObjectWriterAdapter) WriteContent(ctx context.Context, out io.Writer) (int64, error) {
	reader, err := this.object.NewReader(ctx, this.client)
	if err != nil {
		return 0, err
	}

	size, err := io.Copy(out, reader)
	if err != nil {
		return 0, err
	}
	return size, err
}

func NewObjectWriterAdapter(client *storage.Client, object *Object) *ObjectWriterAdapter {
	return &ObjectWriterAdapter{
		client: client,
		object: object,
	}
}
