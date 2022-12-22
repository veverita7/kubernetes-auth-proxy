package main

import (
	"log"
	"net/url"
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/veverita7/kubernetes-auth-proxy/pkg/middlewares/authenticator"
	"github.com/veverita7/kubernetes-auth-proxy/pkg/middlewares/impersonator"
	"github.com/veverita7/kubernetes-auth-proxy/pkg/proxy"
)

func main() {
	log.Fatal(run())
}

func run() error {
	cfg, err := newKubeConfig()
	if err != nil {
		return err
	}

	url, err := url.Parse(cfg.Host)
	if err != nil {
		return err
	}

	return proxy.New(url).AddMiddlewares(authenticator.New(), impersonator.New()).Run(":8080")
}

func newKubeConfig() (*rest.Config, error) {
	path := filepath.Join(homeDir(), ".kube", "config")
	cfg, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		cfg, err = rest.InClusterConfig()
	}
	return cfg, err
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}
