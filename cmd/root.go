package cmd

import (
	"github.com/spf13/cobra"
	"go-getting-started/cmd/api"
	"os"
)

func init() {
	rootCmd.AddCommand(api.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "go-getting-started",
	Short: "go-getting-started",
	Long:  "go-getting-started",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
