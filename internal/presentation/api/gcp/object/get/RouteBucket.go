package get

import (
	"bytes"
	"github.com/gorilla/mux"
	coreObject "github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	"net/http"
)

type RouteBucket struct {
	Route
}

func (this RouteBucket) Provide(r *mux.Router, h Operation) {
	r.
		HandleFunc("/{bucket:[0-9a-zA-Z-_.]+}/{object:.+}", func(w http.ResponseWriter, req *http.Request) {
			Handle(w, req, this, h)
		}).
		Headers("X-KOMMOL-STRATEGY", "GCP_BUCKET").
		Methods("GET")
}

func (this RouteBucket) ExtractObject(req *http.Request) (coreObject.Object, error) {
	params := mux.Vars(req)
	var b bytes.Buffer

	b.WriteString("gs://")
	b.WriteString(params["bucket"])
	b.WriteString("/")
	b.WriteString(params["object"])

	return gcp.ParseObjectUrlString(b.String())
}
