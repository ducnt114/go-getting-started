package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-getting-started/router"
)

var Cmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Long:  `api`,
	Run: func(cmd *cobra.Command, args []string) {
		startApi()
	},
}

func startApi() {
	r, err := router.InitRouter()
	if err != nil {
		panic(err)
	}
	fmt.Println("start api server at :8080")
	_ = r.Run(":8080")
}
