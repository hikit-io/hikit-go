package hfgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type Controller interface {
	GroupName() string
	//Middlewares : prefix is global pre-middleware, suffix is global post middleware
	Middlewares() (prefix, suffix []gin.HandlerFunc)
}

type HandleFunc func() (method, routeUri, version string, handlerFuncs []gin.HandlerFunc)
type MethodHandleFunc func() (routeUri, version string, handlerFuncs []gin.HandlerFunc)
type MethodSnakeHandleFunc func() (version string, handlerFuncs []gin.HandlerFunc)

type NewHandleFunc func() MethodHandleFunc

var (
	_EmptyHandleFunc = HandleFunc(
		func() (method, routeUri, version string, handlerFunc []gin.HandlerFunc) {
			return "", "", "", nil
		},
	)
	_EmptyMethodHandleFunc = MethodHandleFunc(
		func() (routeUri, version string, handlerFunc []gin.HandlerFunc) {
			return "", "", nil
		},
	)
	_EmptyMethodSnakeHandleFunc = MethodSnakeHandleFunc(
		func() (version string, handlerFunc []gin.HandlerFunc) {
			return "", nil
		},
	)
	THandleFunc            = reflect.TypeOf(_EmptyHandleFunc)
	TMethodHandleFunc      = reflect.TypeOf(_EmptyMethodHandleFunc)
	TMethodSnakeHandleFunc = reflect.TypeOf(_EmptyMethodSnakeHandleFunc)
)

func MatchMethod(methodName string) string {
	methodName = strings.ToUpper(methodName)
	switch {
	case strings.HasPrefix(methodName, http.MethodGet):
		return http.MethodGet
	case strings.HasPrefix(methodName, http.MethodPost):
		return http.MethodPost
	case strings.HasPrefix(methodName, http.MethodPut):
		return http.MethodPut
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
				case v.Method(i).Type().ConvertibleTo(THandleFunc):
					{
						f := v.Method(i).Convert(THandleFunc).Interface().(HandleFunc)
						httpMethod, routeUri, version, handlerFunc := f()
						switch {
						case version != "" && httpMethod != "":
							r.Group(version).Group(c.GroupName(), prefix...).Handle(httpMethod, routeUri, append(handlerFunc, suffix...)...)
							continue
						case version != "" && httpMethod == "":
							r.Group(version).Group(c.GroupName(), prefix...).Any(routeUri, append(handlerFunc, suffix...)...)
							continue
						}
					}
				case v.Method(i).Type().ConvertibleTo(TMethodHandleFunc):
					{
						f := v.Method(i).Convert(TMethodHandleFunc).Interface().(MethodHandleFunc)
						routeUri, version, handlerFunc := f()
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
				case v.Method(i).Type().ConvertibleTo(TMethodSnakeHandleFunc):
					{
						f := v.Method(i).Convert(TMethodSnakeHandleFunc).Interface().(MethodSnakeHandleFunc)
						version, handlerFunc := f()
						methodName := v.Type().Method(i).Name
						httpMethod := MatchMethod(methodName)
						routeUri := MatchUrl(methodName[len(httpMethod):])
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
	r1, _ := regexp.Compile("By[A-Z]")
	r2, _ := regexp.Compile("[A-Z]")
	name = r1.ReplaceAllStringFunc(name, func(s string) string {
		ss := []rune(s)
		ss[0] = '/'
		ss[1] = ':'
		ss[2] = ss[2] - ('A' - 'a')
		return string(ss)
	})
	name = r2.ReplaceAllStringFunc(name, func(s string) string {
		ss := []rune(s)
		ss = append(ss, ss[0]-('A'-'a'))
		ss[0] = '/'
		return string(ss)
	})

	return string(name)
}
