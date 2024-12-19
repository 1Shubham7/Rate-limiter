package main

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Client struct {
	limiter *rate.Limiter
	lastSeen time.Time
}

var (
	mu sync.Mutex
	clients = make(map[string]*Client)
)

func perClientRateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _ , err := net.SplitHostPort(r.RemoteAddr) // gets the IP address of client
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, found := clients[ip]; !found {
			clients[ip] = &Client{
				limiter: rate.NewLimiter(2,5),
				lastSeen: time.Now(),
			}
		}

		mu.Lock()

		c := clients[ip]
		c.lastSeen = time.Now()
		if !c.limiter.Allow() {
			mu.Unlock()
			message := Message {
				Status: "429",
				Body: "Too many requests, please try again later",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(message)
			return
		}

		mu.Unlock()
		next(w, r)
	})
}