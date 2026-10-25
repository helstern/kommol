package app

import (
	"github.com/helstern/kommol/internal/core/http"
	"io"
)

type HttpObject struct {
	Headers []http.Header
	Body    io.Reader
	// Status is the HTTP status to answer with; 0 means 200.
	Status int
}
