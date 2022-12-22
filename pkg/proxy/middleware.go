package proxy

import "net/http"

type Middleware interface {
	Handler(h http.Handler) http.Handler
}
