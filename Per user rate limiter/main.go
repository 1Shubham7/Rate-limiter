package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body  string `json:"body"`
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	message  := &Message{
		Status: "200",
		Body:  "Hey you have hit the endpoint",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("Server with Per Client Rate Limiter running on port 7777")
	http.ListenAndServe(":7777", nil)

	http.Handle("/api", perClientRateLimiter(endpointHandler))
}