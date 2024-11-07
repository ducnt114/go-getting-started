package kafka_producer

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "kafka_producer",
	Short: "kafka_producer",
	Long:  `kafka_producer`,
	Run: func(cmd *cobra.Command, args []string) {
		startKafkaProducer()
	},
}

func startKafkaProducer() {
	//injector := do.New()
	//defer func() {
	//	_ = injector.Shutdown()
	//}()
	//conf.Inject(injector)
	//cf := do.MustInvoke[*conf.Config](injector)

	saramaConf := sarama.NewConfig()
	saramaConf.Version = sarama.V1_1_0_0
	saramaConf.Producer.Flush.Messages = 1
	saramaConf.Producer.Flush.Frequency = 1 * time.Second
	saramaConf.Producer.Return.Successes = true
	saramaConf.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(
		[]string{}, saramaConf)
	if err != nil {
		panic(err)
	}

	msg := &MessageExchange{
		Event: "order_create",
		Data:  "order_id_123456",
	}
	msgBytes, _ := json.Marshal(msg)

	producerMsg := &sarama.ProducerMessage{
		Topic:   "topic_1",
		Key:     nil,
		Value:   sarama.ByteEncoder(msgBytes),
		Headers: nil,
	}
	partition, offset, err := producer.SendMessage(producerMsg)
	fmt.Println("partition:", partition, "offset:", offset, "err:", err)
	//producer.Input() <- producerMsg

	time.Sleep(3 * time.Second)
	fmt.Println("produce message success")
}

type MessageExchange struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
