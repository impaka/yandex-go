package main

import (
	"context"
	"fmt"
	"net/http"
)

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		for _, ch := range name {
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				continue
			} else {
				fmt.Fprint(w, "hello dirty hacker")
				return
			}
		}

		next.ServeHTTP(w, r)
	}

}

type ctxKey string

const nameKey ctxKey = "name"

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			name = "stranger"
		}
		newCTX := context.WithValue(r.Context(), nameKey, name)
		rCopy := r.WithContext(newCTX)
		next.ServeHTTP(w, rCopy)
	}

}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(nameKey)
	nameStr, ok := name.(string)
	if !ok {
		return
	}
	input := "hello " + nameStr
	w.Write([]byte(input))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", SetDefaultName(Sanitize(HelloHandler)))
	http.ListenAndServe(":8080", mux)
}
