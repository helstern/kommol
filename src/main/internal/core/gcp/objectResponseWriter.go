package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/infrastructure/http/headers"
	"io"
	"net/http"
)

type ObjectWriter struct {
	object *storage.ObjectHandle
}

func (this *ObjectWriter) ModifyHeaders(ctx context.Context, headerMap http.Header) error {
	attr, err := this.object.Attrs(ctx)
	if err != nil {
		return err
	}

	if attr.ContentType != "" {
		headerMap.Set(headers.ContentType, attr.ContentType)
	}

	if attr.ContentLanguage != "" {
		headerMap.Set(headers.ContentLanguage, attr.ContentLanguage)
	}

	if attr.ContentEncoding != "" {
		headerMap.Set(headers.ContentEncoding, attr.ContentEncoding)
	}

	if attr.ContentDisposition != "" {
		headerMap.Set(headers.ContentDisposition, attr.ContentDisposition)
	}

	if attr.Size != 0 {
		headerMap.Set(headers.ContentLength, headers.FormatInt(attr.Size))
	}

	if attr.CacheControl != "" {
		headerMap.Set(headers.CacheControl, attr.CacheControl)
	}
	return nil
}

func (this *ObjectWriter) WriteContent(ctx context.Context, out io.Writer) (int64, error) {
	reader, err := this.object.NewReader(ctx)
	if err != nil {
		return 0, err
	}

	size, err := io.Copy(out, reader)
	if err != nil {
		return 0, err
	}
	return size, err
}

func NewObjectWriter(object *storage.ObjectHandle) object.Writer {
	return &ObjectWriter{
		object: object,
	}
}
