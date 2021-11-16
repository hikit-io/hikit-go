package hkgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var TraceKey = "trace_id"

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Err        error       `json:"err,omitempty"`
	Code       TCode       `json:"code,omitempty"`
	TraceId    string      `json:"trace_id,omitempty"`
	statusCode StatusCode
	replace    Replace
}

type Option func(response *Response)

type TCode int

const (
	OK TCode = 1024 << iota
	ErrCode
)

func Ok(c *gin.Context, data interface{}, opts ...Option) {
	res := &Response{
		Code:       OK,
		Data:       data,
		statusCode: http.StatusOK,
		TraceId:    c.GetString(TraceKey),
	}
	for _, opt := range opts {
		opt(res)
	}
	if res.replace != nil {
		c.JSON(int(res.statusCode), res.replace())
		return
	}
	c.JSON(int(res.statusCode), res)
}

func Err(c *gin.Context, opts ...Option) {
	res := &Response{
		statusCode: http.StatusOK,
		Code:       ErrCode,
		TraceId:    c.GetString(TraceKey),
	}
	for _, op := range opts {
		op(res)
	}
	if res.replace != nil {
		c.JSON(int(res.statusCode), res.replace())
		return
	}
	c.JSON(int(res.statusCode), res)
}

type StatusCode int

func WithStatusCode(code StatusCode) Option {
	return func(r *Response) {
		r.statusCode = code
	}
}

func WithData(data interface{}) Option {
	return func(r *Response) {
		r.Data = data
	}
}

func WithErr(err error) Option {
	return func(r *Response) {
		r.Err = err
	}
}

func WithMsg(msg string) Option {
	return func(r *Response) {
		r.Msg = msg
	}
}

type Replace func() interface{}

func ReplaceAll(data interface{}) Option {
	return func(r *Response) {
		r.replace = func() interface{} {
			return data
		}
	}
}
