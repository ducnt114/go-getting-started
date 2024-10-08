package kafka_consumer

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "kafka-consumer",
	Short: "kafka-consumer",
	Long:  `kafka-consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		startKafkaConsumer()
	},
}

func startKafkaConsumer() {
	fmt.Println("start kafka consumer...")
}
