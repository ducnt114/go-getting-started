package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"go-getting-started/dto"
	"go-getting-started/utils"
	"net/http"
	"regexp"
	"strings"
)

var ignoreAuthEndpoint = []string{
	`/api/v1/auth/login`,
}

func ignoreAuth(c *gin.Context) bool {
	for _, v := range ignoreAuthEndpoint {
		if match, _ := regexp.MatchString(v, c.Request.URL.Path); match {
			return true
		}
	}
	return false
}

func Auth(di *do.Injector) gin.HandlerFunc {
	jwtUtil := do.MustInvoke[utils.JWTUtil](di)

	return func(ctx *gin.Context) {
		if ignoreAuth(ctx) {
			ctx.Next()
			return
		}
		header := ctx.Request.Header.Get("Authorization")
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		parts := strings.Split(header, " ")
		if len(parts) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		switch parts[0] {
		case "Basic":
			// TODO basic auth
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		case "Bearer":
			var userClaim dto.JwtClaims
			err := jwtUtil.ParseClaims(parts[1], &userClaim)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.Set("user_uuid", userClaim.UserUUID)
			ctx.Set("user_name", userClaim.UserName)
		default:
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
