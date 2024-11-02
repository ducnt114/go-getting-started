package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"go-getting-started/log"
	"go-getting-started/utils"
	"net/http"
)

func Authorization(di *do.Injector) gin.HandlerFunc {
	enforcer := do.MustInvoke[*utils.Enforcer](di)

	return func(ctx *gin.Context) {
		user := ctx.GetString("user_name")
		method := ctx.Request.Method
		path := ctx.Request.URL.Path

		result, _ := enforcer.E.Enforce(user, path, method)
		if result {
			ctx.Next()
		} else {
			log.Warnw(ctx, "casbin check failed", "user", user, "path", path, "method", method)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
