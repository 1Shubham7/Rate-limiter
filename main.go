package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	
)

type Message struct{
	Status string `json:"status"`
	Body string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request){
	message := &Message{
		Status: "200",
		Body: "Hello World",
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	// encodes the message into a json object and writes it to the writer
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func main(){
	fmt.Println("Server with Token Bucket Rate Limiter running on port 6666")
	http.Handle("/api", RateLimiter(endpointHandler))
	err := http.ListenAndServe(":6666", nil)
	if err != nil {
		panic(err)
	}
}
