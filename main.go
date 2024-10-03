package main

import (
	"github.com/beego/beego/v2/server/web"
	"go-getting-started/router"
)

func main() {
	router.InitRouter()

	web.Run()
}
