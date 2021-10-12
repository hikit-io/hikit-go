# hfgin

## Feature

## Example 
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hfunc/hfunc-go/hfutils/hfgin"
)

type _Controller struct{}

func (_ _Controller) GroupName() string {
	return "hfunc"
}

func (_ _Controller) Middlewares() ([]gin.HandlerFunc, []gin.HandlerFunc) {
	return []gin.HandlerFunc{
			func(c *gin.Context) {
				fmt.Println("global prefix")
				c.Next()
			},
		}, []gin.HandlerFunc{
			func(c *gin.Context) {
				fmt.Println("global subfix")
				c.Next()
			},
		}
}

func (_ _Controller) Version() string {
	return "v1"
}

func (_ _Controller) GETPing() (routeUri, version string, handlerFuncs []gin.HandlerFunc) {
	return "ping", "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			c.JSON(200, "pong")
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func (_ _Controller) GETHelloHfunc_id() ( version string, handlerFuncs []gin.HandlerFunc) {
	return  "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			c.JSON(200, "pong")
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func (_ _Controller) GETHelloHfunc() ( version string, handlerFuncs []gin.HandlerFunc) {
	return  "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			c.JSON(200, "pong")
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func main() {
	r := gin.Default()
	hfgin.RegisterController(r, &_Controller{})
	r.Run(":8081")
}
```
result:
```
[GIN-debug] GET    /v1/hfunc/hello/hfunc     --> main._Controller.Middlewares.func2 (6 handlers)
[GIN-debug] GET    /v1/hfunc/hello/hfunc/:id --> main._Controller.Middlewares.func2 (6 handlers)
[GIN-debug] GET    /v1/hfunc/ping            --> main._Controller.Middlewares.func2 (6 handlers)
```