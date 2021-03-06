package hkgin

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"go.hikit.io/hkctx"
)

func Tracer(traceName string) gin.HandlerFunc {
	TraceKey = traceName
	return func(c *gin.Context) {
		c.Set(TraceKey, uuid.New().String())
		c.Next()
	}
}

type logger interface {
	Info(ctx hkctx.Ctx, msg string, keyAndValues ...interface{})
}

func Logger(l logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		c.Next()

		l.Info(c, "logger",
			zap.String("method", c.Request.Method),
			zap.String("host", c.Request.Host),
			zap.String("cost_time", time.Now().Sub(now).String()),
			zap.String("uri", c.Request.RequestURI),
		)
	}
}

func Paramer(l logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodGet {
			params := map[string][]string{}
			c.ShouldBind(&params)
			l.Info(c, "logger",
				zap.Any("params", params),
			)
			c.Next()
			return
		}
		switch c.ContentType() {
		case binding.MIMEJSON:
			params := map[string]interface{}{}
			c.ShouldBind(&params)
			l.Info(c, "logger",
				zap.Any("params", params),
			)
		case binding.MIMEXML, binding.MIMEXML2:
		case binding.MIMEPROTOBUF:
		case binding.MIMEMSGPACK, binding.MIMEMSGPACK2:
		case binding.MIMEYAML:
		case binding.MIMEMultipartPOSTForm:
		default: // case MIMEPOSTForm:
		}
		c.Next()

	}
}
