package gcp

import (
	"testing"

	"cloud.google.com/go/storage"
	"github.com/helstern/kommol/internal/infrastructure/gcp"
)

func TestBucketNotFoundPageWinsOverTheDefault(t *testing.T) {
	if got := gcp.NotFoundPage(&storage.BucketWebsite{NotFoundPage: "missing.html"}, "404.html"); got != "missing.html" {
		t.Fatalf("got %q", got)
	}
}

func TestDefaultNotFoundPageWhenTheBucketSetsNone(t *testing.T) {
	if got := gcp.NotFoundPage(&storage.BucketWebsite{}, "404.html"); got != "404.html" {
		t.Fatalf("got %q", got)
	}
	if got := gcp.NotFoundPage(nil, ""); got != "" {
		t.Fatalf("got %q, want none", got)
	}
}
