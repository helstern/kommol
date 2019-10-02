package routes

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object"
	"net/http"
)

type bucket struct{}

func (b bucket) Provide(r *mux.Router, h object.GetHandler) {
	r.
		HandleFunc("/{bucket:[0-9a-zA-Z-_.]+}/{object:.+}", func(w http.ResponseWriter, req *http.Request) {
			b.Handle(w, req, h)
		}).
		Headers("X-KOMMOL-STRATEGY", "GCP_BUCKET").
		Methods("GET")
}

func (_ bucket) ExtractPath(req *http.Request) string {
	params := mux.Vars(req)
	var b bytes.Buffer

	b.WriteString("gs://")
	b.WriteString(params["bucket"])
	b.WriteString("/")
	b.WriteString(params["object"])

	return b.String()
}

func (b bucket) Handle(w http.ResponseWriter, req *http.Request, h object.GetHandler) {
	path := b.ExtractPath(req)
	h(w, path)
}

var (
	routeBucket *bucket
)

func Bucket() *bucket {
	if routeBucket == nil {
		routeBucket = &bucket{}
	}
	return routeBucket
}
