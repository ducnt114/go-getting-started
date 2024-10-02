package controller

import (
	"github.com/gin-gonic/gin"
	"go-getting-started/dto"
	"net/http"
)

type UserController struct {
}

func (c *UserController) GetInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &dto.User{
		Name: "duc",
		Age:  32,
	})
}
