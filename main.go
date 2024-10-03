package main

import (
	"fmt"
	"go-getting-started/router"
)

func main() {
	r := router.InitRouter()
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
