package impersonator

import (
	"fmt"
	"net/http"
	"time"
)

func New() Impersonator {
	return &impersonator{}
}

type Impersonator interface {
	Handler(h http.Handler) http.Handler
}

type impersonator struct{}

func (*impersonator) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: create impersonation logic
		fmt.Printf("[%s] impersonator middleware test log\n", time.Now().Format(time.RFC3339))
		h.ServeHTTP(w, r)
	})
}
