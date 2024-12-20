package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	toolbooth "github.com/didip/tollbooth/v8"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	message := &Message{
		Status: "200",
		Body:   "Hello World",
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	// encodes the message into a json object and writes it to the writer
	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func main() {
	message := Message {
		Status: "200",
		Body: "Hello World",
	}
	msg, err := json.Marshal(message)
	if err != nil {
		fmt.Print(err)
		return
	}

	// 1 is no. or req/sec allowed, nil is reqest handler called when limit is reached
	limiter := toolbooth.NewLimiter(1, nil)
	limiter.SetMessageContentType("application/json")
	limiter.SetMessage(string(msg))
	http.Handle("/api", toolbooth.LimitFuncHandler(limiter, endpointHandler))

	err = http.ListenAndServe(":6666", nil)
	if err != nil {
		fmt.Print(err)
		return
	}
}