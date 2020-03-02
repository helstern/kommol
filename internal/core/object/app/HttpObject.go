package app

import (
	"github.com/helstern/kommol/internal/core/http"
	"io"
)

type HttpObject struct {
	Headers []http.Header
	Body    io.Reader
}
