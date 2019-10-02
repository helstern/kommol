package object

import "net/http"

type GetHandler func(w http.ResponseWriter, path string)
