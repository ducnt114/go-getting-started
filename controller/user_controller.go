package controller

import (
	"github.com/gin-gonic/gin"
	"go-getting-started/dto"
	"go-getting-started/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

// GetUserByGender get user by gender
// @BasePath /user
// @Summary get user by gender
// @Description get user by gender
// @Accept json
// @Produce json
// @Param header-1 header string true "valid-value"
// @Param gender query string true "gender description"
// @Success 200 {object} dto.UserResponse
// @Router /user [get]
func (u *UserController) GetUserByGender(ctx *gin.Context) {
	gender := ctx.Query("gender")
	//requestID, existed := ctx.Get("request_id")
	//if existed {
	//}
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
