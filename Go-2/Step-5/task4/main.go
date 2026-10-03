package main

import (
	"io"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	fetchHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := request()
		w.Write(data)
	})
	http.Handle("/readSource", http.TimeoutHandler(fetchHandler, maxTimeout, ""))
	http.ListenAndServe(":8080", nil)
}

func request() []byte {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8081/provideData", nil)
	if err != nil {
		return nil
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return body
}
