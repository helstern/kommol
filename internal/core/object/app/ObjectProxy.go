package app

import (
	"context"
	"github.com/helstern/kommol/internal/core"
	"github.com/helstern/kommol/internal/core/object"
	"reflect"
)

type ObjectProxy struct {
	resolvers map[core.CloudProviders]ObjectProvider
}

func (this *ObjectProxy) Http(ctx context.Context, o object.Object) (*HttpObject, error) {

	resolver := this.resolvers[o.Provider]
	container, err := resolver.WebsiteContainer(ctx, o.Path)
	if err != nil {
		return nil, err
	}

	var actualObject object.Object
	resolvedObject, err := container.ResolveObject(o.Path)
	if err != nil {
		return nil, err
	}

	if reflect.DeepEqual(resolvedObject, o) {
		actualObject = o
	} else {
		actualObject = resolvedObject
	}

	headers, err := resolver.Headers(ctx, actualObject.Path)
	if err != nil {
		return nil, err
	}

	data, err := resolver.Data(ctx, actualObject.Path)
	if err != nil {
		return nil, err
	}

	return &HttpObject{
		Headers: headers,
		Body:    data,
	}, nil
}

func NewObjectProxy(resolvers map[core.CloudProviders]ObjectProvider) ObjectProxy {
	return ObjectProxy{resolvers: resolvers}
}
