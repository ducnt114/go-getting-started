package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-getting-started/dto"
	"go-getting-started/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func (u *UserController) GetUserByGender(ctx *gin.Context) {
	gender := ctx.Query("gender")
	val, existed := ctx.Get("key-1")
	if existed {
		fmt.Println("key-1 value: ", val)
	}
	resp := u.UserService.GetUserByGender(gender)

	ctx.JSON(http.StatusOK, resp)
}

func (u *UserController) ListUser(ctx *gin.Context) {
	name := ctx.Query("name")
	age := ctx.Query("age")
	ctx.JSON(http.StatusOK, &dto.UserResponse{
		Name: name,
		Age:  age,
	})
}
