package driver_api

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:   "driver_api",
	Short: "driver_api",
	Long:  `driver_api`,
	Run: func(cmd *cobra.Command, args []string) {
		startApi()
	},
}

func startApi() {
	fmt.Println("driver api")
	http.HandleFunc("/api/v1/driver", driverHandler)
	addr := ":8080"
	fmt.Println("Server is running on ", addr)
	_ = http.ListenAndServe(addr, nil)
}

type Order struct {
	ID string `json:"id"`
}

func driverHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req Order
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.ID == "invalid-order-id" {
			fmt.Println("fail to looking for driver for order: ", req.ID)
			http.Error(w, "fail", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(req)
	case http.MethodDelete:
		//id := r.URL.Query().Get("id")
		//if id == "" {
		//	http.Error(w, "Missing driver ID", http.StatusBadRequest)
		//	return
		//}
		//if _, exists := drivers[id]; !exists {
		//	http.Error(w, "Driver not found", http.StatusNotFound)
		//	return
		//}
		//delete(drivers, id)
		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
