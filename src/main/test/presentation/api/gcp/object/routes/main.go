package routes

import "net/http"

func objectTestHandler(w http.ResponseWriter, path string) {
	_, err := w.Write([]byte(path))
	if err != nil {
		panic(err)
	}
}
