package cmd

import (
	"github.com/spf13/cobra"
	"go-getting-started/cmd/api"
	"go-getting-started/cmd/data_api"
	"go-getting-started/cmd/demo_error"
	"go-getting-started/cmd/event_sourcing"
	"go-getting-started/cmd/kafka_consumer"
	"go-getting-started/cmd/migrate"
	"go-getting-started/cmd/saga/driver_api"
	"go-getting-started/cmd/saga/order_api"
	"go-getting-started/cmd/saga/payment_api"
	"os"
)

func init() {
	rootCmd.AddCommand(api.Cmd)
	rootCmd.AddCommand(migrate.Cmd)
	rootCmd.AddCommand(kafka_consumer.Cmd)
	rootCmd.AddCommand(demo_error.Cmd)
	rootCmd.AddCommand(data_api.Cmd)
	rootCmd.AddCommand(order_api.Cmd)
	rootCmd.AddCommand(payment_api.Cmd)
	rootCmd.AddCommand(driver_api.Cmd)
	rootCmd.AddCommand(event_sourcing.Cmd)
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
