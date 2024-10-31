package cmd

import (
	"github.com/spf13/cobra"
	"go-getting-started/cmd/api"
	"go-getting-started/cmd/data_api"
	"go-getting-started/cmd/demo_error"
	"go-getting-started/cmd/kafka_consumer"
	"go-getting-started/cmd/migrate"
	"os"
)

func init() {
	rootCmd.AddCommand(api.Cmd)
	rootCmd.AddCommand(migrate.Cmd)
	rootCmd.AddCommand(kafka_consumer.Cmd)
	rootCmd.AddCommand(demo_error.Cmd)
	rootCmd.AddCommand(data_api.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "go-getting-started",
	Short: "go-getting-started",
	Long:  "go-getting-started",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
