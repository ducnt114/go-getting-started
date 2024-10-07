package main

import (
	"go-getting-started/conf"
	"go-getting-started/router"
)

func init() {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	r, err := router.InitRouter()
	if err != nil {
		panic(err)
	}
	_ = r.Run(":8080")
}
