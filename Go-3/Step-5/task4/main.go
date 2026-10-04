package main

import (
	"context"
	"fmt"
	"net/http"
)

type ctxKey string

const nameKey ctxKey = "username"

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if username == "" || password == "" || !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized\n"))
			return
		}
		newCTX := context.WithValue(r.Context(), nameKey, username)
		rCopy := r.WithContext(newCTX)
		next.ServeHTTP(w, rCopy)

	}

}
func answerHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(nameKey)
	usernameStr, ok := username.(string)
	if !ok {
		return
	}

	fmt.Fprintf(w, "Welcome, %s!", usernameStr)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/answer/", Authorization(answerHandler))
	http.ListenAndServe(":8080", mux)
}
