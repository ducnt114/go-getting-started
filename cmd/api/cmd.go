package api

import (
	"context"
	"github.com/spf13/cobra"
	"go-getting-started/pkg/log"
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
	log.Infow(context.Background(), "Starting API")
}
