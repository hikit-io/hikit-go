package hkhttp

import (
	"net/http"

	"golang.org/x/oauth2"
)

type TokenHandle func(req *http.Request, token *oauth2.Token)

func WithUrlToken(name string) TokenHandle {
	return func(req *http.Request, token *oauth2.Token) {
		req.URL.Query().Set(name, token.AccessToken)
	}
}

func WithHeaderToken(name string) TokenHandle {
	return func(req *http.Request, token *oauth2.Token) {
		req.Header.Set(name, token.AccessToken)
	}
}

type Transport struct {
	Base        http.RoundTripper
	Source      oauth2.TokenSource
	TokenHandle TokenHandle
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqBodyClosed := false
	if req.Body != nil {
		defer func() {
			if !reqBodyClosed {
				req.Body.Close()
			}
		}()
	}

	if t.Source != nil {
		token, err := t.Source.Token()
		if err != nil {
			return nil, err
		}
		if t.TokenHandle != nil {
			t.TokenHandle(req, token)
		} else {
			token.SetAuthHeader(req)
		}
	}

	return t.RoundTrip(req)
}
