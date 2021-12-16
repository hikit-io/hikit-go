package hkhttp

import (
	"net"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var (
	DefaultRoundTripper = http.Transport{
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
	}
)

type clientOptions struct {
	trans http.RoundTripper
	oauth oauth2.Config
}

type ClientOption interface {
	apply(options *clientOptions)
}

func NewClient(opts ...ClientOption) *http.Client {
	opt := &clientOptions{
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
		o.apply(opt)
	}
	return &http.Client{
		Transport: opt.trans,
	}
}

type RoundTripper struct {
	http.RoundTripper
}

func (r RoundTripper) apply(options *clientOptions) {
	options.trans = r
}

func WithRoundTripper(rt http.RoundTripper) ClientOption {
	return RoundTripper{rt}
}
