package proxy

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var httpTransport *http.Transport

func init() {
	httpTransport = http.DefaultTransport.(*http.Transport).Clone()
	httpTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func New(u *url.URL) Proxy {
	return &proxy{
		server: &httputil.ReverseProxy{
			Director: func(r *http.Request) {
				r.URL.Scheme = u.Scheme
				r.URL.Host = u.Host
				r.URL.Path = u.Path
				r.Header.Add("X-Forwarded-For", r.RemoteAddr)
				r.Header.Add("X-Forwarded-Host", r.Host)
				r.Header.Add("X-Forwarded-Proto", r.Proto)
			},
			Transport: httpTransport,
		},
	}
}

type Proxy interface {
	AddMiddlewares(m ...Middleware) Proxy
	Run(addr string) error
}

type proxy struct {
	server      http.Handler
	middlewares []Middleware
}

func (p *proxy) AddMiddlewares(m ...Middleware) Proxy {
	p.middlewares = append(p.middlewares, m...)
	return p
}

func (p *proxy) Run(addr string) error {
	h := p.server
	for i := range p.middlewares {
		h = p.middlewares[len(p.middlewares)-i-1].Handler(h)
	}
	return http.ListenAndServe(addr, h)
}
