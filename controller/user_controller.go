package controller

import (
	"github.com/gin-gonic/gin"
	"go-getting-started/dto"
	"net/http"
)

type UserController struct {
}

func (u *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, &dto.UserResponse{
		ID: id,
	})
}

func (u *UserController) ListUser(ctx *gin.Context) {
	name := ctx.Query("name")
	age := ctx.Query("age")
	ctx.JSON(http.StatusOK, &dto.UserResponse{
		Name: name,
		Age:  age,
	})
}
