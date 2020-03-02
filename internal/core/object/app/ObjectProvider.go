package app

import (
	"context"
	"github.com/helstern/kommol/internal/core/http"
	"github.com/helstern/kommol/internal/core/object"
	"io"
)

type ObjectProvider interface {
	WebsiteContainer(ctx context.Context, objectPath []string) (object.WebsiteContainer, error)

	Headers(ctx context.Context, objectPath []string) ([]http.Header, error)

	Data(ctx context.Context, objectPath []string) (io.Reader, error)
}
