package gcp

import (
	"github.com/helstern/kommol/internal/core"
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	"reflect"
	"testing"
)

func TestCreateEmptyWebsiteContainer(t *testing.T) {
	container := gcp.EmptyWebsiteContainer()
	if container.Provider != core.GCP {
		//t.Errorf("could not parse path actual: %s, expected %s", actualPath, expectedPath)
		t.Errorf("unexpected provider for an empty website container")
	}
}

func TestCreateNormalWebsiteContainer(t *testing.T) {
	container, err := gcp.NewWebsiteContainer("bucket", "index.html")

	if err != nil {
		t.Errorf("unexpected error returned when gcp.NewWebsiteContainer called with valid arguments")
	}

	if container.Provider != core.GCP {
		//t.Errorf("could not parse path actual: %s, expected %s", actualPath, expectedPath)
		t.Errorf("unexpected provider for an empty website container")
	}

	object, resolvedFileErr := container.ResolveObject([]string{"bucket", "/"})
	if resolvedFileErr != nil {
		t.Errorf("unexpected error returned resolving an object")
	}

	if !reflect.DeepEqual(object.Path, []string{"bucket", "index.html"}) {
		t.Errorf("unexpected resolved object")
	}
}
