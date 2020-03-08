package api

import "net/http"

func RequestId(req *http.Request) string {
	id := req.Header.Get("X-Request-ID")
	return id
}
