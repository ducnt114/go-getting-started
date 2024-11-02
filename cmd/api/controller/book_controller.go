package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"go-getting-started/service"
	"net/http"
)

type BookController interface {
	List(ctx *gin.Context)
	GetData(ctx *gin.Context)
}

type bookCtl struct {
	bookService service.BookService
}

func NewBookController(di *do.Injector) BookController {
	return &bookCtl{
		bookService: do.MustInvoke[service.BookService](di),
	}
}

func (c *bookCtl) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "success")
}

func (c *bookCtl) GetData(ctx *gin.Context) {

	err := c.bookService.GetData(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "success")
}
