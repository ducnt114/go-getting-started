package middlewares

import (
	"crypto/rsa"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/do"
	"go-getting-started/conf"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var ignoreAuthEndpoint = []string{
	`/auth/login`,
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
	cf := do.MustInvoke[*conf.Config](di)
	jwtPubKey, err := os.ReadFile(cf.JWT.PublicKeyFilePath)
	if err != nil {
		panic(err)
	}
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(jwtPubKey)
	if err != nil {
		panic(err)
	}

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
			claims, err := parseAccessToken(parts[1], pubKey)
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.Set("user_uuid", claims.UserUUID)
			ctx.Set("user_name", claims.UserName)
		default:
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}

type jwtClaims struct {
	jwt.RegisteredClaims
	UserUUID string `json:"user_uuid"`
	UserName string `json:"user_name"`
}

func parseAccessToken(accessToken string, key *rsa.PublicKey) (*jwtClaims, error) {
	claims := &jwtClaims{}
	_, err := jwt.NewParser().ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return claims, err
}
