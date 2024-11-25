package payment_api

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "payment_api",
	Short: "payment_api",
	Long:  `payment_api`,
	Run: func(cmd *cobra.Command, args []string) {
		startApi()
	},
}

func startApi() {
	fmt.Println("payment api")
	http.HandleFunc("/api/v1/payment", paymentHandler)
	http.HandleFunc("/api/v1/refund", refundHandler)
	addr := ":8081"
	fmt.Println("Server is running on", addr)
	http.ListenAndServe(addr, nil)
}

type Order struct {
	ID string `json:"id"`
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Order
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Payment for order: ", req.ID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func refundHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Order
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Refund for order: ", req.ID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}
