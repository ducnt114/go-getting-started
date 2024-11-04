package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "mqtt",
	Short: "mqtt",
	Long:  `mqtt`,
	Run: func(cmd *cobra.Command, args []string) {
		startMqtt()
	},
}

const (
	broker = "tcp://localhost:1883"
	topic  = "test/topic"
	//clientID = "go_mqtt_client"
)

func startMqtt() {
	// Define the MQTT message handler for incoming messages
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message on topic %s: %s\n", msg.Topic(), msg.Payload())
	}

	clientID := fmt.Sprintf("go_mqtt_client_%d", time.Now().Unix())

	// Set up MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(messageHandler)

	// Create the MQTT client
	client := mqtt.NewClient(opts)

	// Connect to the MQTT broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Error connecting to broker: %v\n", token.Error())
		os.Exit(1)
	}
	fmt.Println("Connected to MQTT broker")

	// Subscribe to a topic
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Printf("Error subscribing to topic: %v\n", token.Error())
		os.Exit(1)
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)

	// Publish a message to the topic
	message := "Hello, MQTT!"
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	fmt.Printf("Published message to topic %s: %s\n", topic, message)

	// Wait a few seconds to receive messages
	time.Sleep(30 * time.Second)

	// Disconnect from the broker
	client.Disconnect(250)
	fmt.Println("Disconnected from MQTT broker")
}
