# hfhttp

## Features

- oauth http client
- custom http auth function
- http3 server

## Usage

### Server

#### http2

#### http3


### Client

#### oauth client

```go
package main

import (
	"github.com/hfunc/hfunc-go/hfhttp"
	"golang.org/x/oauth2"
)

func main() {
	cli := hfhttp.NewClient(hfhttp.WithRoundTripper(&hfhttp.Transport{
		Base: nil,
		Source: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "123",
		}),
	}))
	cli.Get("http://127.0.0.1")
}
```
Final:
```http request
GET http://127.0.0.1
Authorization: Basic 123
```
#### In URL

```go
package main

import (
	"github.com/hfunc/hfunc-go/hfhttp"
	"golang.org/x/oauth2"
)

func main() {
	cli := hfhttp.NewClient(hfhttp.WithRoundTripper(&hfhttp.Transport{
		Base: nil,
		Source: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "123",
		}),
		TokenHandle: hfhttp.WithUrlToken("api_key"),
	}))
	cli.Get("http://127.0.0.1")
}
```
```http request
GET http://127.0.0.1
```
#### In Header

```go
package main

import (
	"github.com/hfunc/hfunc-go/hfhttp"
	"golang.org/x/oauth2"
)

func main() {
	cli := hfhttp.NewClient(hfhttp.WithRoundTripper(&hfhttp.Transport{
		Base: nil,
		Source: oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: "123",
		}),
		TokenHandle: hfhttp.WithHeaderToken("api_key"),
	}))
	cli.Get("http://127.0.0.1")
}
```
```http request
GET http://127.0.0.1
api_key: 123
```