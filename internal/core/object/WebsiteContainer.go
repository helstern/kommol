package object

import (
	"github.com/helstern/kommol/internal/core"
	"github.com/pkg/errors"
	"reflect"
)

type WebsiteContainer struct {
	Provider     core.CloudProviders
	path         []string
	indexFile    string
	notFoundPage string
}

// NotFoundObject is the page served, with status 404, for an object that does not exist; false when there is none.
func (this *WebsiteContainer) NotFoundObject() (Object, bool) {
	if this.notFoundPage == "" {
		return Object{}, false
	}
	objectPath := append(append([]string{}, this.path...), this.notFoundPage)
	return Object{Provider: this.Provider, Path: objectPath}, true
}

func (this *WebsiteContainer) ResolveObject(path []string) (Object, error) {

	container := path[0:len(this.path)]
	if !reflect.DeepEqual(this.path, container) {
		return Object{
			Provider: this.Provider,
			Path:     path,
		}, errors.New("object does not belong in this container")
	}

	objectKey := path[len(this.path):]
	objectPath := append(container[0:0], container...)

	if len(objectKey) == 0 && this.indexFile != "" {
		objectPath = append(objectPath, this.indexFile)
	} else if len(objectKey) == 1 && (objectKey[0] == "" || objectKey[0] == "/") && this.indexFile != "" {
		objectPath = append(objectPath, this.indexFile)
	} else if len(objectKey) > 0 {
		objectPath = append(objectPath, objectKey...)
	}

	return Object{
		Provider: this.Provider,
		Path:     objectPath,
	}, nil

}

func NewWebsiteContainer(provider core.CloudProviders, path []string, indexFile string) (WebsiteContainer, error) {
	return NewWebsiteContainerWithNotFound(provider, path, indexFile, "")
}

func NewWebsiteContainerWithNotFound(provider core.CloudProviders, path []string, indexFile string, notFoundPage string) (WebsiteContainer, error) {
	if len(path) == 0 {
		return WebsiteContainer{
			Provider:  provider,
			path:      nil,
			indexFile: "",
		}, errors.New("path must not be empty")
	}

	return WebsiteContainer{
		Provider:     provider,
		path:         path,
		indexFile:    indexFile,
		notFoundPage: notFoundPage,
	}, nil
}
