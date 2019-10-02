package routes

import (
	"github.com/gorilla/mux"
	"github.com/helstern/kommol/internal/presentation/api/gcp/object"
)

func Get(r *mux.Router, h object.GetHandler) {
	Bucket().Provide(r, h)
	Website().Provide(r, h)
}
