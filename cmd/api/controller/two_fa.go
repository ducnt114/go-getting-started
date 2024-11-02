package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"go-getting-started/dto"
	"go-getting-started/log"
	"go-getting-started/service"
	"net/http"
)

type TwoFaController interface {
	Get2Fa(*gin.Context)
	Setup2Fa(*gin.Context)
}

type twoFaCtl struct {
	authService service.AuthService
}

func NewTwoFaController(di *do.Injector) TwoFaController {
	return &twoFaCtl{
		authService: do.MustInvoke[service.AuthService](di),
	}
}

func (c *twoFaCtl) Get2Fa(ctx *gin.Context) {
	userID, exist := ctx.Get("user_id")
	if !exist {
		log.Errorw(ctx, "error when get user_id from token")
		ctx.JSON(http.StatusInternalServerError, "error when get user_id from token")
		return
	}
	resp, err := c.authService.GetTwoFa(ctx, userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *twoFaCtl) Setup2Fa(ctx *gin.Context) {
	req := &dto.SetupTwoFaRequest{}
	_ = ctx.ShouldBind(req)
	resp, err := c.authService.SetupTwoFa(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, resp)
}
