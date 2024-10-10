package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.New().String()
			c.Request.Header.Set("X-Request-Id", requestId)
		}
		c.Set("request_id", requestId)
		c.Next()
	}
}
