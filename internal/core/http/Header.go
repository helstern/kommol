package http

import (
	"net/http"
	"strconv"
	"time"
)

type Header struct {
	Name  string
	Value string
}

func NewIntHeader(name string, value int64) Header {
	return Header{
		Name:  name,
		Value: strconv.FormatInt(value, 10),
	}
}

func NewTimeHeader(name string, value time.Time) Header {
	return Header{
		Name:  name,
		Value: value.UTC().Format(http.TimeFormat),
	}
}

func NewContentTypeHeader(value string) Header {
	return Header{
		Name:  "content-type",
		Value: value,
	}
}

func NewContentLengthHeader(value int64) Header {
	return NewIntHeader("content-length", value)
}

func NewContentEncodingHeader(value string) Header {
	return Header{
		Name:  "content-encoding",
		Value: value,
	}
}

func NewContentDispositionHeader(value string) Header {
	return Header{
		Name:  "content-disposition",
		Value: value,
	}
}

func NewContentLanguageHeader(value string) Header {
	return Header{
		Name:  "content-language",
		Value: value,
	}
}

func NewCacheControlHeader(value string) Header {
	return Header{
		Name:  "cache-control",
		Value: value,
	}
}
