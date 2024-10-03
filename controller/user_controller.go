package controller

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

func (u *UserController) Get() {
	u.Ctx.WriteString("hello world from user controller")
}

type CreateUserResponse struct {
	Username string `json:"username"`
}

func (u *UserController) Post() {
	_ = u.Ctx.JSONResp(&CreateUserResponse{
		Username: "duc",
	})
}
