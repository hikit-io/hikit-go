package rtm

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type Client struct {
	httpCli        *http.Client
	appId          string
	appCertificate string
	expire         uint64
}

type Options struct {
	AppId string
	Token string
}

type Option interface {
	Apply(*Options)
}

type Callback func(*Options)

func Ctx() context.Context {
	return context.Background()
}

func New(opts ...Option) *Client {
	opt := &Options{}
	for _, o := range opts {
		o.Apply(opt)
	}
	authCli := oauth2.NewClient(
		Ctx(),
		oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: opt.Token,
		}),
	)
	cli := &Client{
		httpCli: authCli,
	}
	return cli
}

type BashAuth struct {
	clientId, clientSecret string
}

func (b BashAuth) Apply(o *Options) {
	o.Token = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", b.clientId, b.clientSecret)))
}

func WithBashAuth(clientId, clientSecret string) Option {
	return BashAuth{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func GenTokenHandler(req *http.Request) {

}

func (c *Client) genToken(rtmUid string) string {
	return ""
}
