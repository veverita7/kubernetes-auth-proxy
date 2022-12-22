package authenticator

import (
	"fmt"
	"net/http"
	"time"
)

func New() Authenticator {
	return &authenticator{}
}

type Authenticator interface {
	Handler(h http.Handler) http.Handler
}

type authenticator struct{}

func (*authenticator) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: create authentication logic
		fmt.Printf("[%s] authentication middleware test log\n", time.Now().Format(time.RFC3339))
		h.ServeHTTP(w, r)
	})
}
