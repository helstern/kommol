package get

import (
	"context"
	"errors"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/helstern/kommol/internal/core"
	"github.com/helstern/kommol/internal/core/http"
	"github.com/helstern/kommol/internal/core/logging/app"
	"github.com/helstern/kommol/internal/core/object"
	objectapp "github.com/helstern/kommol/internal/core/object/app"
	"github.com/helstern/kommol/internal/infrastructure/logging"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object/get"
)

type provider struct{ objects map[string]string; notFound string }

func (p provider) WebsiteContainer(ctx context.Context, path []string) (object.WebsiteContainer, error) {
	return object.NewWebsiteContainerWithNotFound(core.GCP, path[0:1], "index.html", p.notFound)
}
func (p provider) body(path []string) (string, error) {
	k := strings.Join(path[1:], "/")
	if k == "broken" {
		return "", errors.New("backend down")
	}
	b, ok := p.objects[k]
	if !ok {
		return "", object.ErrNotFound
	}
	return b, nil
}
func (p provider) Headers(ctx context.Context, path []string) ([]http.Header, error) {
	_, err := p.body(path)
	return []http.Header{}, err
}
func (p provider) Data(ctx context.Context, path []string) (io.Reader, error) {
	b, err := p.body(path)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(b), nil
}

func serve(p provider, key string) *httptest.ResponseRecorder {
	px := objectapp.NewObjectProxy(map[core.CloudProviders]objectapp.ObjectProvider{core.GCP: p})
	op := get.NewOperation(px, app.LoggerFactory(logging.NewLoggerWithFields))
	w := httptest.NewRecorder()
	op(w, httptest.NewRequest("GET", "/"+key, nil), object.Object{Provider: core.GCP, Path: []string{"bucket", key}})
	return w
}


func TestMissingObjectWithPageAnswers404WithThePage(t *testing.T) {
	w := serve(provider{objects: map[string]string{"404.html": "gone"}, notFound: "404.html"}, "nope")
	if w.Code != 404 || w.Body.String() != "gone" {
		t.Fatalf("got %d %q", w.Code, w.Body.String())
	}
}

func TestMissingObjectWithoutPageAnswersPlain404(t *testing.T) {
	w := serve(provider{objects: map[string]string{}}, "nope")
	if w.Code != 404 {
		t.Fatalf("got %d, want 404", w.Code)
	}
}

func TestOtherErrorsAnswer500(t *testing.T) {
	w := serve(provider{objects: map[string]string{}}, "broken")
	if w.Code != 500 {
		t.Fatalf("got %d, want 500", w.Code)
	}
}

func TestFoundObjectAnswers200(t *testing.T) {
	w := serve(provider{objects: map[string]string{"a.html": "A"}}, "a.html")
	if w.Code != 200 || w.Body.String() != "A" {
		t.Fatalf("got %d %q", w.Code, w.Body.String())
	}
}
