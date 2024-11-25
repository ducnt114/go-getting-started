package order_api

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "order_api",
	Short: "order_api",
	Long:  `order_api`,
	Run: func(cmd *cobra.Command, args []string) {
		startApi()
	},
}

func startApi() {
	fmt.Println("order api")
	http.HandleFunc("/api/v1/order", orderHandler)
	addr := ":8082"
	fmt.Println("Server is running on", addr)
	_ = http.ListenAndServe(addr, nil)
}

type Order struct {
	ID string `json:"id"`
}

var orders = make(map[string]Order)

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	orders[order.ID] = order

	// SAGA
	saga := &Saga{}
	saga.AddStep(&PaymentStep{})
	saga.AddStep(&DriverStep{})

	sagaErr := saga.Execute(order)
	if sagaErr != nil {
		fmt.Println("error when execute saga, detail: ", sagaErr)
		http.Error(w, "fail", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
