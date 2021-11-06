package hkhttp

import (
	"crypto/tls"
	"net/http"
	"time"
)

type serverOptions struct {
	network      string
	addr         string
	handler      http.Handler
	tlsConfig    *tls.Config
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
}

type ServerOption interface {
	apply(server *serverOptions)
}

type addr string

func (a addr) apply(o *serverOptions) {
	o.addr = string(a)
}

func WithAddr(addr addr) ServerOption {
	return addr
}

type handler struct {
	http.Handler
}

func (h handler) apply(o *serverOptions) {
	o.handler = h
}

func WithHandler(h http.Handler) ServerOption {
	return handler{h}
}

type tlsConfig tls.Config

func (t *tlsConfig) apply(o *serverOptions) {
	o.tlsConfig = (*tls.Config)(t)
}

// WithTLSConfig with hserver tls config.
func WithTLSConfig(conf *tls.Config) ServerOption {
	return (*tlsConfig)(conf)
}
