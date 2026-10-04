package main

import (
	"fmt"
	"net/http"
	"sync"
)

var query int

type ctxKey string

var a, b int = 0, 1
var mu sync.Mutex

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		query++
		mu.Unlock()
		next.ServeHTTP(w, r)

	}

}
func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	result := a
	a, b = b, a+b
	mu.Unlock()
	fmt.Fprint(w, result)

}
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", query)
	mu.Unlock()
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", Metrics(FibonacciHandler))
	mux.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(":8080", mux)
}
