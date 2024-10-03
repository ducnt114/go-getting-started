package controller

import "github.com/gin-gonic/gin"

type PingController struct {
}

type Pong struct {
	Message string `json:"message_id"`
}

func (p *PingController) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
