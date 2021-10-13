# hfgin

## Feature

## Example 
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hfunc/hfunc-go/hflog"
	"github.com/hfunc/hfunc-go/hfutils/hfgin"
	"go.uber.org/zap"
)

type _Controller struct{}

func (_ _Controller) GroupName() string {
	return "hfunc"
}

func (_ _Controller) Middlewares() (prefix, suffix []gin.HandlerFunc) {
	return []gin.HandlerFunc{
			hfgin.TraceId("trace_id"),
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

func (_ _Controller) GetUserById() (version string, handlerFuncs []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			params := c.Param("id")
			hflog.Info(c, "GetUserById", zap.String("id", params))
			hfgin.Ok(c, params)
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func (_ _Controller) PostUserById() (version string, handlerFuncs []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			hfgin.Ok(c, "pong")
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func (_ _Controller) PutUserById() (version string, handlerFuncs []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			hfgin.Ok(c, "pong")
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func (_ _Controller) DeleteUserById() (version string, handlerFuncs []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			hfgin.Ok(c, "pong")
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func (_ _Controller) PatchUserById() (version string, handlerFuncs []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			hfgin.Ok(c, "pong")
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func (_ _Controller) PostUserByIdName() (version string, handlerFuncs []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			hfgin.Ok(c, "pong")
		},
		func(c *gin.Context) {
			fmt.Println("subfix")
			c.Next()
		},
	}
}

func (_ _Controller) GetUserList() (version string, handlerFuncs []gin.HandlerFunc) {
	return "v1", []gin.HandlerFunc{
		func(c *gin.Context) {
			hfgin.Ok(c, "pong")
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
[GIN-debug] DELETE /v1/hfunc/user/:id        --> main._Controller.Middlewares.func2 (7 handlers)
[GIN-debug] GET    /v1/hfunc/user/:id        --> main._Controller.Middlewares.func2 (7 handlers)
[GIN-debug] GET    /v1/hfunc/user/list       --> main._Controller.Middlewares.func2 (7 handlers)
[GIN-debug] PATCH  /v1/hfunc/user/:id        --> main._Controller.Middlewares.func2 (7 handlers)
[GIN-debug] POST   /v1/hfunc/user/:id        --> main._Controller.Middlewares.func2 (7 handlers)
[GIN-debug] POST   /v1/hfunc/user/:id/name   --> main._Controller.Middlewares.func2 (7 handlers)
[GIN-debug] PUT    /v1/hfunc/user/:id        --> main._Controller.Middlewares.func2 (7 handlers)
```