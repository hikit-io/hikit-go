package hfhttp

import (
	"net"
	"net/http"
	"time"
)

type options struct {
	trans http.RoundTripper
}

type Options func(*options)

func NewClient(opts ...Options) *http.Client {
	opt := &options{
		trans: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	for _, o := range opts {
		o(opt)
	}
	return &http.Client{
		Transport: opt.trans,
	}
}

func WithRoundTripper(rt http.RoundTripper) Options {
	return func(opt *options) {
		opt.trans = rt
	}
}
