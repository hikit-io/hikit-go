package hfgin

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceId(traceName string) gin.HandlerFunc {
	TraceKey = traceName
	return func(c *gin.Context) {
		c.Set(TraceKey, uuid.New().String())
		c.Next()
	}
}
