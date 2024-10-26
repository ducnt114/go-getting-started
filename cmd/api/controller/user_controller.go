package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"go-getting-started/dto"
	"go-getting-started/service"
	"net/http"
	"strconv"
)

type UserController interface {
	GetUserById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userCtl struct {
	userService service.UserService
}

func NewUserController(di *do.Injector) UserController {
	return &userCtl{
		userService: do.MustInvoke[service.UserService](di),
	}
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
func (c *userCtl) GetUserById(ctx *gin.Context) {
	userID := ctx.Param("id")
	uid, _ := strconv.ParseInt(userID, 10, 64)

	resp, err := c.userService.GetUserById(ctx, uint(uid))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
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
func (c *userCtl) Create(ctx *gin.Context) {
	req := &dto.CreateUserReq{}
	_ = ctx.ShouldBind(req)
	resp, err := c.userService.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, resp)
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
func (c *userCtl) Delete(ctx *gin.Context) {
	userID := ctx.Param("id")
	uid, _ := strconv.ParseInt(userID, 10, 64)
	ctx.JSON(http.StatusOK, &dto.User{
		ID:   uint(uid),
		Name: "duc",
		Age:  32,
	})
}

func (c *userCtl) List(ctx *gin.Context) {
	name := ctx.Query("name")

	resp, err := c.userService.List(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *userCtl) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	resp, err := c.userService.Update(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
