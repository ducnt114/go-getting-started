package cmd

import (
	"github.com/spf13/cobra"
	"go-getting-started/cmd/api"
	"go-getting-started/cmd/data_api"
	"go-getting-started/cmd/demo_error"
	"go-getting-started/cmd/kafka_consumer"
	"go-getting-started/cmd/kafka_producer"
	"go-getting-started/cmd/migrate"
	"go-getting-started/cmd/mqtt"
	"go-getting-started/cmd/rabbitmq"
	"go-getting-started/cmd/rabbitmq_consumer"
	"go-getting-started/cmd/sse"
	"go-getting-started/cmd/websocket_server"
	"os"
)

func init() {
	rootCmd.AddCommand(api.Cmd)
	rootCmd.AddCommand(migrate.Cmd)
	rootCmd.AddCommand(demo_error.Cmd)
	rootCmd.AddCommand(data_api.Cmd)
	rootCmd.AddCommand(websocket_server.Cmd)
	rootCmd.AddCommand(sse.Cmd)
	rootCmd.AddCommand(mqtt.Cmd)
	rootCmd.AddCommand(rabbitmq.Cmd)
	rootCmd.AddCommand(rabbitmq_consumer.Cmd)
	rootCmd.AddCommand(kafka_producer.Cmd)
	rootCmd.AddCommand(kafka_consumer.Cmd)
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
