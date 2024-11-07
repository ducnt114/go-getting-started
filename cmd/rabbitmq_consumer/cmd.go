package rabbitmq_consumer

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
	"log"
)

var Cmd = &cobra.Command{
	Use:   "rabbitmq_consumer",
	Short: "rabbitmq_consumer",
	Long:  `rabbitmq_consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		startRabbitmqConsumer()
	},
}

func startRabbitmqConsumer() {
	user := "dev"
	pass := ""
	host := ""
	port := 5672
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/", user, pass, host, port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")
	queueName := "direct_queue_2"
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
