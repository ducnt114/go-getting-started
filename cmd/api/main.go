package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Routes
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/divide", divide)

	// Start server on port 8080
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ping(w http.ResponseWriter, r *http.Request) {
	type pong struct {
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&pong{Message: "ok"})
}

type divideInput struct {
	Val1 int `json:"val_1"`
	Val2 int `json:"val_2"`
}

type divideOutput struct {
	Result int `json:"result"`
}

func divide(w http.ResponseWriter, r *http.Request) {
	var input divideInput
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.Unmarshal(reqBytes, &input); err != nil {
		fmt.Println(err)
		return
	}
	var res = &divideOutput{
		Result: calc(input.Val1, input.Val2),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(res)
}

func calc(a, b int) int {
	res := a / b
	return res
}
