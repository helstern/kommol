package routes

import (
	"github.com/gorilla/mux"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebsiteSuccess(t *testing.T) {

	absoluteUrlRequest := httptest.NewRequest(
		"GET",
		"http://www.mozilla.com/bucket/object",
		nil,
	)
	absoluteUrlRequest.Header.Add("X-KOMMOL-STRATEGY", "GCP_WEBSITE")

	relativeUrlRequest := httptest.NewRequest(
		"GET",
		"/bucket/object",
		nil,
	)
	relativeUrlRequest.Header.Add("Host", "www.mozilla.com")
	//simulate the result
	relativeUrlRequest.URL.Host = ""
	relativeUrlRequest.Host = "www.mozilla.com"

	reqs := []*http.Request{
		absoluteUrlRequest,
		relativeUrlRequest,
	}
	for _, req := range reqs {
		req.Header.Add("X-KOMMOL-STRATEGY", "GCP_WEBSITE")
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
