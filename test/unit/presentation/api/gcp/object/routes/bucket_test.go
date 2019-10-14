package routes

import (
	"github.com/gorilla/mux"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object/routes"
	"net/http/httptest"
	"testing"
)

func TestBucketSuccess(t *testing.T) {

	router := mux.NewRouter()
	routes.Bucket().Provide(router, objectTestHandler)

	actualUrl := "http://www.example.com/bucket/object"
	expectedPath := "gs://bucket/object"

	req := httptest.NewRequest(
		"GET",
		actualUrl,
		nil,
	)
	req.Header.Add("X-KOMMOL-STRATEGY", "GCP_BUCKET")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	actualPath := rr.Body.String()
	if actualPath != expectedPath {
		t.Logf("actual path: %s", actualPath)
		t.Errorf("the bucket route does not handle expected urls")
	}
}
