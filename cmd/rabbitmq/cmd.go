package rabbitmq

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
)

var Cmd = &cobra.Command{
	Use:   "rabbitmq",
	Short: "rabbitmq",
	Long:  `rabbitmq`,
	Run: func(cmd *cobra.Command, args []string) {
		startRabbitmqDemo()
	},
}

func startRabbitmqDemo() {
	// "amqp://guest:guest@localhost:5672/"
	amqpURL := fmt.Sprintf("amqp://%v:%v@%v:%v/",
		"", "", "", "")

	// Step 1: Establish a connection to RabbitMQ
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Step 2: Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Step 3: Declare an exchange (type: direct)
	exchangeName := "exch_headers"
	//err = ch.ExchangeDeclare(
	//	exchangeName, // exchange name
	//	"direct",     // exchange type
	//	true,         // durable
	//	false,        // auto-deleted
	//	false,        // internal
	//	false,        // no-wait
	//	nil,          // arguments
	//)
	//if err != nil {
	//	log.Fatalf("Failed to declare an exchange: %v", err)
	//}

	// Step 4: Publish a message to the exchange
	routingKey := ""
	message := "Hello, RabbitMQ!"

	header := make(amqp.Table)
	header["format"] = "json"
	header["event"] = "create"

	err = ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			Headers:     header,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf("Message sent: %s", message)
}
