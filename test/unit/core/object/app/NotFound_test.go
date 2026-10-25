package app

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/helstern/kommol/internal/core"
	"github.com/helstern/kommol/internal/core/http"
	"github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/core/object/app"
)

// fakeProvider serves the objects it holds; any other key is object.ErrNotFound, and "broken" fails otherwise.
type fakeProvider struct {
	objects      map[string]string
	notFoundPage string
}

func (f fakeProvider) WebsiteContainer(ctx context.Context, path []string) (object.WebsiteContainer, error) {
	return object.NewWebsiteContainerWithNotFound(core.GCP, path[0:1], "index.html", f.notFoundPage)
}

func (f fakeProvider) find(path []string) (string, error) {
	key := strings.Join(path[1:], "/")
	if key == "broken" {
		return "", errors.New("backend down")
	}
	body, ok := f.objects[key]
	if !ok {
		return "", object.ErrNotFound
	}
	return body, nil
}

func (f fakeProvider) Headers(ctx context.Context, path []string) ([]http.Header, error) {
	_, err := f.find(path)
	return []http.Header{}, err
}

func (f fakeProvider) Data(ctx context.Context, path []string) (io.Reader, error) {
	body, err := f.find(path)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(body), nil
}

func proxy(p fakeProvider) app.ObjectProxy {
	return app.NewObjectProxy(map[core.CloudProviders]app.ObjectProvider{core.GCP: p})
}

func get(t *testing.T, p fakeProvider, key string) (*app.HttpObject, string, error) {
	px := proxy(p)
	res, err := px.Http(context.Background(), object.Object{Provider: core.GCP, Path: []string{"bucket", key}})
	if err != nil {
		return nil, "", err
	}
	body, _ := ioutil.ReadAll(res.Body)
	return res, string(body), nil
}

func TestFoundObjectIsServedWithoutStatus(t *testing.T) {
	res, body, err := get(t, fakeProvider{objects: map[string]string{"a.html": "A"}, notFoundPage: "404.html"}, "a.html")
	if err != nil || body != "A" || res.Status != 0 {
		t.Fatalf("got status %v body %q err %v", res, body, err)
	}
}

func TestMissingObjectServesTheNotFoundPageWith404(t *testing.T) {
	res, body, err := get(t, fakeProvider{objects: map[string]string{"404.html": "gone"}, notFoundPage: "404.html"}, "nope")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if res.Status != 404 || body != "gone" {
		t.Fatalf("got status %d body %q, want 404 \"gone\"", res.Status, body)
	}
}

func TestMissingObjectWithoutNotFoundPageIsErrNotFound(t *testing.T) {
	_, _, err := get(t, fakeProvider{objects: map[string]string{}}, "nope")
	if !errors.Is(err, object.ErrNotFound) {
		t.Fatalf("got %v, want object.ErrNotFound", err)
	}
}

func TestMissingNotFoundPageIsErrNotFound(t *testing.T) {
	_, _, err := get(t, fakeProvider{objects: map[string]string{}, notFoundPage: "404.html"}, "nope")
	if !errors.Is(err, object.ErrNotFound) {
		t.Fatalf("got %v, want object.ErrNotFound", err)
	}
}

func TestOtherErrorsAreNotNotFound(t *testing.T) {
	_, _, err := get(t, fakeProvider{objects: map[string]string{"404.html": "gone"}, notFoundPage: "404.html"}, "broken")
	if err == nil || errors.Is(err, object.ErrNotFound) {
		t.Fatalf("got %v, want a non-not-found error", err)
	}
}
