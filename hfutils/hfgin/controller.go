package hfgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type Controller interface {
	GroupName() string
	Middlewares() (prefix, suffix []gin.HandlerFunc)
}

type MethodHandleFunc func() (routeUri, version string, handlerFuncs []gin.HandlerFunc)
type MethodSnakeHandleFunc func() (version string, handlerFuncs []gin.HandlerFunc)

type NewHandleFunc func() MethodHandleFunc

var (
	_EmptyMethodHandleFunc = MethodHandleFunc(
		func() (routeUri, version string, handlerFunc []gin.HandlerFunc) {
			return "", "", nil
		},
	)
	_EmptyNewMethodHandleFunc = NewHandleFunc(
		func() MethodHandleFunc {
			return func() (routeUri, version string, handlerFunc []gin.HandlerFunc) {
				return "", "", nil
			}
		},
	)
	_EmptyMethodSnakeHandleFunc = MethodSnakeHandleFunc(
		func() (version string, handlerFunc []gin.HandlerFunc) {
			return "", nil
		},
	)
	TMethodHandleFunc         = reflect.TypeOf(_EmptyMethodHandleFunc)
	TNewMethodHandleFunc      = reflect.TypeOf(_EmptyNewMethodHandleFunc)
	TNewMethodSnakeHandleFunc = reflect.TypeOf(_EmptyMethodSnakeHandleFunc)
)

//func RegisterHandle(r gin.IRouter, handleFuncs ...MethodHandleFunc) {
//	for _, handleFunc := range handleFuncs {
//		 url, version, handles := handleFunc()
//		r.Group(version).Handle(method, url, handles...)
//	}
//}

func MatchMethod(methodName string) string {
	switch {
	case strings.HasPrefix(methodName, http.MethodGet):
		return http.MethodGet
	case strings.HasPrefix(methodName, http.MethodPost):
		return http.MethodPost
	case strings.HasPrefix(methodName, http.MethodPatch):
		return http.MethodPatch
	case strings.HasPrefix(methodName, http.MethodHead):
		return http.MethodHead
	case strings.HasPrefix(methodName, http.MethodDelete):
		return http.MethodDelete
	case strings.HasPrefix(methodName, http.MethodConnect):
		return http.MethodConnect
	case strings.HasPrefix(methodName, http.MethodOptions):
		return http.MethodOptions
	case strings.HasPrefix(methodName, http.MethodTrace):
		return http.MethodTrace
	}
	return ""
}

func RegisterController(r gin.IRouter, controllers ...Controller) {
	if r != nil {
		for _, c := range controllers {
			v := reflect.ValueOf(c)
			prefix, suffix := c.Middlewares()
			for i := 0; i < v.NumMethod(); i++ {
				switch {
				case v.Method(i).Type().ConvertibleTo(TNewMethodHandleFunc):
					{
						outs := v.Method(i).Call(nil)
						f := outs[0].Interface().(MethodHandleFunc)
						routeUri, version, handlerFunc := f()
						methodName := v.Type().Method(i).Name
						httpMethod := MatchMethod(methodName)
						switch {
						case version != "" && httpMethod != "":
							r.Group(version).Group(c.GroupName(), prefix...).Handle(httpMethod, routeUri, append(handlerFunc, suffix...)...).Use()
							continue
						case version != "" && httpMethod == "":
							r.Group(version).Group(c.GroupName(), prefix...).Any(routeUri, append(handlerFunc, suffix...)...)
							continue
						}
					}
				case v.Method(i).Type().ConvertibleTo(TMethodHandleFunc):
					{
						outs := v.Method(i).Call(nil)
						routeUri, version, handlerFunc := outs[0].Interface().(string),
							outs[1].Interface().(string),
							outs[2].Interface().([]gin.HandlerFunc)
						methodName := v.Type().Method(i).Name
						httpMethod := MatchMethod(methodName)
						switch {
						case version != "" && httpMethod != "":
							r.Group(version).Group(c.GroupName(), prefix...).Handle(httpMethod, routeUri, append(handlerFunc, suffix...)...)
							continue
						case version != "" && httpMethod == "":
							r.Group(version).Group(c.GroupName(), prefix...).Any(routeUri, append(handlerFunc, suffix...)...)
							continue
						}
					}
				case v.Method(i).Type().ConvertibleTo(TNewMethodSnakeHandleFunc):
					{
						outs := v.Method(i).Call(nil)
						version, handlerFunc := outs[0].Interface().(string),
							outs[1].Interface().([]gin.HandlerFunc)
						methodName := v.Type().Method(i).Name
						httpMethod := MatchMethod(methodName)
						routeUri := MatchUrl(strings.TrimPrefix(methodName, httpMethod))
						switch {
						case version != "" && httpMethod != "":
							r.Group(version).Group(c.GroupName(), prefix...).Handle(httpMethod, routeUri, append(handlerFunc, suffix...)...).Use()
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

func MatchUrl(name string) string {
	newstr := make([]rune, 0)
	for idx, chr := range name {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newstr = append(newstr, '/')
			}
			chr -= 'A' - 'a'
		}

		if chr == '_' {
			chr = '/'
			newstr = append(newstr, chr, ':')
			continue
		}
		newstr = append(newstr, chr)
	}
	return string(newstr)
}
