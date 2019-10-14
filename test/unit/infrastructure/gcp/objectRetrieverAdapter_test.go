package gcp

import (
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	"testing"
)

type stub struct {
	Object *gcp.Object
}

func (o *stub) RetrieveObject(obj *gcp.Object) (object.Writer, error) {
	o.Object = obj
	return nil, nil
}

func TestObjectRetrieverAdapterSuccess(t *testing.T) {

	var stub = &stub{}
	var retriever = gcp.NewObjectRetrieverAdapter(stub)

	var expectedPath = "gs://bucket/key"
	_, _ = retriever.Retrieve(expectedPath)
	var actualPath = stub.Object.GetPath()

	if actualPath != expectedPath {
		t.Errorf("could not parse path actual: %s, expected %s", actualPath, expectedPath)
	}
}
