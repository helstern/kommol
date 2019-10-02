package object

import (
	"context"
	"io"
	"net/http"
)

type Writer interface {
	ModifyHeaders(ctx context.Context, headers http.Header) error
	WriteContent(ctx context.Context, out io.Writer) (int64, error)
}
