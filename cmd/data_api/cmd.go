package data_api

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var Cmd = &cobra.Command{
	Use:   "data_api",
	Short: "data_api",
	Long:  `data_api`,
	Run: func(cmd *cobra.Command, args []string) {
		startDataApi()
	},
}

func startDataApi() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		response := Response{Message: "Hello, World!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	fmt.Println("Starting data api server on :8081")
	_ = http.ListenAndServe(":8081", nil)
}

type Response struct {
	Message string `json:"message"`
}
