package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Store struct {
	users  map[int]User
	nextID int
	mu     sync.RWMutex
}
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewStore() *Store {
	return &Store{
		users:  map[int]User{},
		nextID: 1,
		mu:     sync.RWMutex{},
	}
}

func (s *Store) CreateUser(name string, age int) User {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.nextID
	s.nextID++
	user := User{
		ID:   id,
		Name: name,
		Age:  age,
	}
	s.users[id] = user
	return user

}

func (s *Store) GetUser(id int) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.users[id]
	if !ok {
		return User{}, false
	}
	return user, true
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(rw, r)
		elapsed := time.Since(start)
		logger.Info("http request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rw.statusCode),
			slog.Duration("duration", elapsed))
	})

}

var store = NewStore()

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}
	user := store.CreateUser(input.Name, input.Age)
	json.NewEncoder(w).Encode(user)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}
	user, ok := store.GetUser(id)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
	json.NewEncoder(w).Encode(user)
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /users/{id}", getUserHandler)
	handler := loggingMiddleware(logger, mux)
	http.ListenAndServe(":8080", handler)
}
