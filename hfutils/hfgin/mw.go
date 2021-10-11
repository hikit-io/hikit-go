package hfgin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hfunc/hfunc-go/hflog"
)

func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(hflog.TraceKey(), uuid.New().String())
		c.Next()
	}
}
