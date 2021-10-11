package hfgin

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	RouterRegister(group *gin.RouterGroup)
	RouterGroupName() string
	Middlewares() []gin.HandlerFunc
	Version() string
}

type HandleFunc func() (httpMethod, routeUri, version string, handlerFuncs []gin.HandlerFunc)
type NewHandleFunc func() HandleFunc

var _emptyHandleFunc = HandleFunc(
	func() (httpMethod, routeUri, version string, handlerFunc []gin.HandlerFunc) {
		return "", "", "", nil
	},
)

var _emptyNewHandleFunc = NewHandleFunc(
	func() HandleFunc {
		return func() (httpMethod, routeUri, version string, handlerFunc []gin.HandlerFunc) {
			return "", "", "", nil
		}
	},
)
