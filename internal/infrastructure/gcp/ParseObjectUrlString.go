package gcp

import (
	"github.com/helstern/kommol/internal/core"
	"github.com/helstern/kommol/internal/core/object"
	"net/url"
	"strings"
)

func ParseObjectUrlString(path string) (object.Object, error) {
	oUrl, err := url.Parse(path)
	if err != nil {
		return object.Object{}, err
	}

	segments := []string{oUrl.Host}

	segments = append(segments, strings.Split(
		strings.Trim(oUrl.Path, "/"),
		"/",
	)...)
	if strings.HasSuffix(oUrl.Path, "/") {
		segments = append(segments, "/")
	}

	return object.Object{
		Path:     segments,
		Provider: core.GCP,
	}, nil
}
