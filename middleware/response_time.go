package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ResponseTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// pre-handler
		startTime := time.Now().UnixMilli()
		fmt.Println("start time: ", startTime)
		fmt.Println("url:", ctx.Request.URL.String())

		header1Val := ctx.Request.Header.Get("header-1")
		// check rate-limit
		//

		if header1Val != "valid-value" {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		ctx.Next()
		//post-handler
		stopTime := time.Now().UnixMilli()
		fmt.Println("exec time: ", stopTime-startTime)
	}
}
