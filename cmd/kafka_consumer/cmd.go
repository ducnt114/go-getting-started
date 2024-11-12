package kafka_consumer

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"log"
)

var Cmd = &cobra.Command{
	Use:   "kafka_consumer",
	Short: "kafka_consumer",
	Long:  `kafka_consumer`,
	Run: func(cmd *cobra.Command, args []string) {
		startKafkaConsumer()
	},
}

func startKafkaConsumer() {
	saramaConf := sarama.NewConfig()
	saramaConf.Version = sarama.V1_1_0_0
	saramaConf.Consumer.Return.Errors = true
	saramaConf.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramaConf.Consumer.Offsets.Initial = sarama.OffsetNewest

	kkClient, err := sarama.NewClient(
		[]string{}, saramaConf)
	if err != nil {
		panic(err)
	}
	kkConsumerGroup, err := sarama.NewConsumerGroupFromClient("group_id_1", kkClient)

	err = kkConsumerGroup.Consume(context.Background(), []string{"topic_1"}, &Consumer{})
	if err != nil {
		fmt.Println("consume error:", err)
	}
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready chan bool
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	//close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
