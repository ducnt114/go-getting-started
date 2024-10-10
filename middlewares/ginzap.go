package middlewares

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go-getting-started/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func GinZap() gin.HandlerFunc {
	return ginzap.GinzapWithConfig(log.Logger(), &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        false,
		SkipPaths:  []string{"/swagger/*"},
		Context: func(c *gin.Context) []zapcore.Field {
			var fields []zapcore.Field
			if requestID := c.Request.Header.Get("X-Request-Id"); requestID != "" {
				fields = append(fields, zap.String("request_id", requestID))
			}
			return fields
		},
	})
}
