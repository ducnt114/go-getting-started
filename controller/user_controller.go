package controller

import (
	"github.com/gin-gonic/gin"
	"go-getting-started/dto"
	"go-getting-started/service"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService service.UserService
}

// GetUserById get user by id
// @BasePath /api/v1
// @Summary get user by id
// @Description get user by id
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} dto.User
// @Router /user/:id [get]
func (c *UserController) GetUserById(ctx *gin.Context) {
	userID := ctx.Param("id")
	uid, _ := strconv.ParseInt(userID, 10, 64)
	resp, err := c.UserService.GetUserById(ctx, uint(uid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, resp)
}

// Create user
// @BasePath /api/v1
// @Summary delete user by id
// @Description delete user by id
// @Accept json
// @Produce json
// @Param user body dto.User true "user info"
// @Success 200 {object} dto.User
// @Router /user [post]
func (c *UserController) Create(ctx *gin.Context) {
	userID := ctx.Param("id")
	uid, _ := strconv.ParseInt(userID, 10, 64)
	ctx.JSON(http.StatusOK, &dto.User{
		ID:   uint(uid),
		Name: "duc",
		Age:  32,
	})
}

// Delete user by id
// @BasePath /api/v1
// @Summary delete user by id
// @Description delete user by id
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} dto.User
// @Router /user/:id [delete]
func (c *UserController) Delete(ctx *gin.Context) {
	userID := ctx.Param("id")
	uid, _ := strconv.ParseInt(userID, 10, 64)
	ctx.JSON(http.StatusOK, &dto.User{
		ID:   uint(uid),
		Name: "duc",
		Age:  32,
	})
}
