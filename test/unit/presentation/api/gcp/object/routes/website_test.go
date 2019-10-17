package routes

import (
	"github.com/gorilla/mux"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func buildRequest(host string, path string, isRelative bool) *http.Request {
	var target string
	if isRelative {
		target = path
	} else {
		target = "http://" + host + path
	}

	request := httptest.NewRequest(
		"GET",
		target,
		nil,
	)
	request.Header.Add("X-KOMMOL-STRATEGY", "GCP_WEBSITE")

	if isRelative {
		request.Header.Add("Host", host)
		//simulate the result
		request.URL.Host = ""
		request.Host = host
	}

	return request
}

func TestWebsiteSuccess(t *testing.T) {
	reqs := []*http.Request{
		buildRequest("www.mozilla.com", "/bucket/object", false),
		buildRequest("www.mozilla.com:80", "/bucket/object", false),
		buildRequest("www.mozilla.com:80", "/bucket/object", true),
		buildRequest("www.mozilla.com", "/bucket/object", true),
	}

	expectedPath := "gs://www.mozilla.com/bucket/object"
	for _, req := range reqs {
		router := mux.NewRouter()
		routes.Website().Provide(router, objectTestHandler)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		actualPath := rr.Body.String()
		if actualPath != expectedPath {
			t.Logf("actual path: %s", actualPath)
			t.Errorf("the website route does not handle expected urls")
		}
	}
}
