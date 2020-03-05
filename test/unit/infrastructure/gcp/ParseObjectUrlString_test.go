package gcp

import (
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	"testing"
)

func TestParseObjectUrlString(t *testing.T) {

	obj, _ := gcp.ParseObjectUrlString("gs://bucket/")
	if len(obj.Path) != 2 {
		t.Errorf("unexpected path length for root object %d", len(obj.Path))
	}

}
