package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"net/http"
)

type BookController interface {
	List(ctx *gin.Context)
}

type bookCtl struct {
}

func NewBookController(di *do.Injector) BookController {
	return &bookCtl{}
}

func (c *bookCtl) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "success")
}
