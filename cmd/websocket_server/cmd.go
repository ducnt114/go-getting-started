package websocket_server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "websocket_server",
	Short: "websocket_server",
	Long:  `websocket_server`,
	Run: func(cmd *cobra.Command, args []string) {
		//startWebsocketServer()
		startSingleChat()
	},
}

func startSingleChat() {
	http.HandleFunc("/ws", handleSingleConnections)

	//go handleMessages()

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func handleSingleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		} else {
			fmt.Println("receive message: ", string(message))
		}
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Println("Write error: ", err)
		} else {
			fmt.Println("send message: ", string(message))
		}
	}
}

func startWebsocketServer() {
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

var clients = make(map[*Client]bool)
var broadcast = make(chan []byte)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v", err)
		return
	}
	defer conn.Close()

	client := &Client{conn: conn, send: make(chan []byte)}
	clients[client] = true

	go handleClient(client)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			delete(clients, client)
			break
		}
		broadcast <- message
	}
}

func handleClient(client *Client) {
	defer client.conn.Close()

	for message := range client.send {
		if err := client.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Write error: %v", err)
			client.conn.Close()
			delete(clients, client)
			break
		}
	}
}

func handleMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(clients, client)
			}
		}
	}
}
