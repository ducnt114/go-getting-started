package main

import "go-getting-started/router"

func init() {

}

func main() {
	r, err := router.InitRouter()
	if err != nil {
		panic(err)
	}
	_ = r.Run(":8080")
}
