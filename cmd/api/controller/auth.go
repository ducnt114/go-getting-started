package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/samber/do"
	"go-getting-started/dto"
	"go-getting-started/log"
	"go-getting-started/service"
	"html/template"
	"net/http"
)

type AuthController interface {
	PasswordLogin(*gin.Context)
	HomePage(*gin.Context)
	SignInWithProvider(*gin.Context)
	CallbackHandler(*gin.Context)
	Success(*gin.Context)
}

type authCtl struct {
	authService service.AuthService
}

func NewAuthController(di *do.Injector) AuthController {
	return &authCtl{
		authService: do.MustInvoke[service.AuthService](di),
	}
}

func (c *authCtl) PasswordLogin(ctx *gin.Context) {
	req := &dto.PasswordLoginRequest{}
	_ = ctx.ShouldBind(req)
	resp, err := c.authService.PasswordLogin(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *authCtl) HomePage(ctx *gin.Context) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(ctx.Writer, gin.H{})
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

func (c *authCtl) SignInWithProvider(ctx *gin.Context) {
	provider := ctx.Param("provider")
	q := ctx.Request.URL.Query()
	q.Add("provider", provider)
	ctx.Request.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (c *authCtl) CallbackHandler(ctx *gin.Context) {
	provider := ctx.Param("provider")
	q := ctx.Request.URL.Query()
	q.Add("provider", provider)
	ctx.Request.URL.RawQuery = q.Encode()

	ggUser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Infow(ctx, "gg user", "user", ggUser)

	ctx.Redirect(http.StatusTemporaryRedirect, "/api/v1/auth/success")
}

func (c *authCtl) Success(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(fmt.Sprintf(`
      <div style="
          background-color: #fff;
          padding: 40px;
          border-radius: 8px;
          box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
          text-align: center;
      ">
          <h1 style="
              color: #333;
              margin-bottom: 20px;
          ">You have Successfull signed in!</h1>
          
          </div>
      </div>
  `)))
}
