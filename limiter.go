package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 5) // so the bucket has 2 tokens added per sec
	// and the maximum number of tokens in the bucket = 5
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			message := Message{
				Status: "429", // too many requests
				Body:   "I ain't coming into your DDOS attack you filthy animal!",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		} else {
			// otherwise we give control to next
			next(w, r)
		}
	})
}
