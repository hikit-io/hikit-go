package hfhttp

import (
	"golang.org/x/oauth2"
	"testing"
)

func TestNewClient(t *testing.T) {
	_ = Transport{
		TokenHandle: WithUrlToken("da"),
	}
	_ = Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "",
		}),
		TokenHandle: WithUrlToken("api_key"),
	}
	_ = Transport{
		Source: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "",
		}),
		TokenHandle: WithHeaderToken("access_token"),
	}
	NewClient(WithRoundTripper(&Transport{
		Base: nil,
		Source: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "123",
		}),
		TokenHandle: WithUrlToken("api_key"),
	}))
}
