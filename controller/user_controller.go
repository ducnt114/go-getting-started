package controller

import (
	"fmt"
	"github.com/getsentry/sentry-go"
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
	sentryCxt := ctx.Request.Context()
	hub := sentry.GetHubFromContext(ctx)
	if hub == nil {
		// Check the concurrency guide for more details: https://docs.sentry.io/platforms/go/concurrency/
		hub = sentry.CurrentHub().Clone()
		sentryCxt = sentry.SetHubOnContext(ctx, hub)
	}
	options := []sentry.SpanOption{
		// Set the OP based on values from https://develop.sentry.dev/sdk/performance/span-operations/
		sentry.WithOpName("http.server"),
		sentry.ContinueFromRequest(ctx.Request),
		sentry.WithTransactionSource(sentry.SourceURL),
	}

	transaction := sentry.StartTransaction(sentryCxt,
		fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL.Path),
		options...,
	)
	defer transaction.Finish()

	//span := sentry.StartSpan(ctx, "UserController.GetUserById")
	//span.Description = "GetUserById_controller"
	//defer span.Finish()

	resp, err := c.UserService.GetUserById(sentryCxt, uint(uid))
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
func (c *UserController) Create(ctx *gin.Context) {
	req := &dto.User{}
	_ = ctx.ShouldBind(req)
	resp, err := c.UserService.CreateUser(ctx, req)
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
func (c *UserController) Delete(ctx *gin.Context) {
	userID := ctx.Param("id")
	uid, _ := strconv.ParseInt(userID, 10, 64)
	ctx.JSON(http.StatusOK, &dto.User{
		ID:   uint(uid),
		Name: "duc",
		Age:  32,
	})
}
