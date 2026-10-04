package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var a, b int = 0, 1
	var query int
	mu := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		result := a
		a, b = b, a+b

		fmt.Fprint(w, result)
		query++
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprint(w, query)

	})

	http.ListenAndServe(":8080", nil)
}
