package hfgin

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type Controller interface {
	GroupName() string
	Middlewares() (prefix, suffix []gin.HandlerFunc)
}

type HandleFunc func() (method, routeUri, version string, handlerFuncs []gin.HandlerFunc)

type NewHandleFunc func() HandleFunc

var _EmptyHandleFunc = HandleFunc(
	func() (method, routeUri, version string, handlerFunc []gin.HandlerFunc) {
		return "", "", "", nil
	},
)

var _EmptyNewHandleFunc = NewHandleFunc(
	func() HandleFunc {
		return func() (method, routeUri, version string, handlerFunc []gin.HandlerFunc) {
			return "", "", "", nil
		}
	},
)

func RegisterHandle(r gin.IRouter, handleFuncs ...HandleFunc) {
	for _, handleFunc := range handleFuncs {
		method, url, version, handles := handleFunc()
		r.Group(version).Handle(method, url, handles...)
	}
}

func RegisterController(r gin.IRouter, controllers ...Controller) {
	if r != nil {
		for _, c := range controllers {
			v := reflect.ValueOf(c)
			prefix, suffix := c.Middlewares()
			for i := 0; i < v.NumMethod(); i++ {
				switch {
				case v.Method(i).Type().ConvertibleTo(reflect.TypeOf(_EmptyNewHandleFunc)):
					{
						outs := v.Method(i).Call(nil)
						f := outs[0].Interface().(HandleFunc)
						httpMethod, routeUri, version, handlerFunc := f()
						switch {
						case version != "" && httpMethod != "":
							r.Group(version).Group(c.GroupName(), prefix...).Handle(httpMethod, routeUri, append(handlerFunc, suffix...)...).Use()
							continue
						case version != "" && httpMethod == "":
							r.Group(version).Group(c.GroupName(), prefix...).Any(routeUri, append(handlerFunc, suffix...)...)
							continue
						}
					}
				case v.Method(i).Type().ConvertibleTo(reflect.TypeOf(_EmptyHandleFunc)):
					{
						outs := v.Method(i).Call(nil)
						httpMethod, routeUri, version, handlerFunc := outs[0].Interface().(string),
							outs[1].Interface().(string),
							outs[2].Interface().(string),
							outs[3].Interface().([]gin.HandlerFunc)
						switch {
						case version != "" && httpMethod != "":
							r.Group(version).Group(c.GroupName(), prefix...).Handle(httpMethod, routeUri, append(handlerFunc, suffix...)...)
							continue
						case version != "" && httpMethod == "":
							r.Group(version).Group(c.GroupName(), prefix...).Any(routeUri, append(handlerFunc, suffix...)...)
							continue
						}
					}
				}
			}
		}
		return
	}
}
