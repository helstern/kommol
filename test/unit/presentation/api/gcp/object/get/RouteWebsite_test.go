package get

import (
	"fmt"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object/get"
	"net/http"
	"net/http/httptest"
	"strings"
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

//func TestWebsiteSuccess(t *testing.T) {
//	reqs := []*http.Request{
//		buildRequest("www.mozilla.com", "/bucket/object", false),
//		buildRequest("www.mozilla.com:80", "/bucket/object", false),
//		buildRequest("www.mozilla.com:80", "/bucket/object", true),
//		buildRequest("www.mozilla.com", "/bucket/object", true),
//	}
//
//	expectedPath := "gs://www.mozilla.com/bucket/object"
//	for _, req := range reqs {
//		route := &get.RouteWebsite{}
//
//		router := mux.NewRouter()
//		route.Provide(router, createGetObjectOperation())
//		rr := httptest.NewRecorder()
//		router.ServeHTTP(rr, req)
//
//		actualPath := rr.Body.String()
//		if actualPath != expectedPath {
//			t.Logf("actual path: %s", actualPath)
//			t.Errorf("the website route does not handle expected urls")
//		}
//	}
//}

func TestExtractObject(t *testing.T) {

	host := "www.mozilla.com"
	path := "/bucket/object"

	reqs := []*http.Request{
		buildRequest(host, path, false),
		buildRequest(fmt.Sprintf("%s:80", host), path, false),
		buildRequest(fmt.Sprintf("%s:80", host), path, true),
		buildRequest(host, path, true),
	}

	for _, req := range reqs {
		route := &get.RouteWebsite{}
		object, _ := route.ExtractObject(req)

		if host+path != strings.Join(object.Path, "/") {
			t.Logf("actual path: %s", strings.Join(object.Path, "/"))
			t.Errorf("unexpected extracted object")
		}
	}
}
