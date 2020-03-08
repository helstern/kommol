package get

import (
	"bytes"
	"github.com/gorilla/mux"
	coreObject "github.com/helstern/kommol/internal/core/object"
	"github.com/helstern/kommol/internal/infrastructure/gcp"
	"net/http"
	"strings"
)

type RouteWebsite struct {
	Route
}

func (this RouteWebsite) ProvideRoot(r *mux.Router, h Operation) RouteWebsite {
	r.
		HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			Handle(w, req, this, h)
		}).
		Headers("X-KOMMOL-STRATEGY", "GCP_WEBSITE").
		Methods("GET")

	return this
}

func (this RouteWebsite) ProvideObject(r *mux.Router, h Operation) RouteWebsite {
	r.
		HandleFunc("/{_:.+}", func(w http.ResponseWriter, req *http.Request) {
			Handle(w, req, this, h)
		}).
		Headers("X-KOMMOL-STRATEGY", "GCP_WEBSITE").
		Methods("GET")

	return this
}

func (this RouteWebsite) Provide(r *mux.Router, h Operation) {
	this.ProvideObject(r, h).ProvideRoot(r, h)
}

func (this RouteWebsite) ExtractObject(req *http.Request) (coreObject.Object, error) {

	bucket := req.URL.Host
	if bucket == "" {
		bucket = req.Host
	}
	bucket = strings.Split(bucket, ":")[0]

	var b bytes.Buffer
	b.WriteString("gs://")
	b.WriteString(bucket)
	b.WriteString(req.URL.Path)

	return gcp.ParseObjectUrlString(b.String())
}
