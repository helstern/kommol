package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	"google.golang.org/api/option"
	"testing"
)

func TestObjectRetrieverAdapterSuccess(t *testing.T) {

	ctx := context.Background()
	client, _ := storage.NewClient(ctx, option.WithCredentialsFile("testdata/credentials.json"))

	var actualPath string
	var objectWriterFactory = gcp.NewObjectWriterFactory(client)

	retriever := gcp.NewObjectRetriever(func(bucket string, object string) object.Writer {
		actualPath = "gs://" + bucket + "/" + object
		return objectWriterFactory(bucket, object)
	})

	var expectedPath = "gs://bucket/key"
	_, _ = retriever.Retrieve(expectedPath)

	if actualPath != expectedPath {
		t.Errorf("could not parse path actual: %s, expected %s", actualPath, expectedPath)
	}
}
