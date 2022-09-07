package hkhttp

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucas-clemente/quic-go/http3"
)

func TestNewServer(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "OK")
	})
	err := Serve(WithAddr(":9090"), WithHandler(r))
	//http3.ConfigureTLSConfig(GenTLSConfigNoErr())

	//err := http3.ListenAndServe(":9090", "", "", r)
	if err != nil {
		panic(err)
	}
}

func TestClient(t *testing.T) {
	cli := http.DefaultClient
	cli.Transport = &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	resp, err := cli.Get("https://localhost:9090/ping")
	fmt.Println(resp, err)
}
