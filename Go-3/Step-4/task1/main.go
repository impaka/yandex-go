package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if len(name) == 0 {
			fmt.Fprint(w, "hello stranger")
			return
		}
		for _, ch := range name {
			if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				continue
			} else {
				fmt.Fprint(w, "hello dirty hacker")
				return
			}
		}
		if len(name) > 0 {
			fmt.Fprintf(w, "hello %s", name)

		}

	})
	http.ListenAndServe(":8080", nil)
}

//  Напишите веб-сервер,
//  который будет возвращать приветствие с именем пользователя,
//  полученным из параметра запроса. Если параметр пустой или
// отсутствует, сервер должен вернуть приветствие "hello stranger".
// Если в ответе не только английские буквы, приветствие должно гласить:
// "hello dirty hacker" (можете применить функцию из пакета strings или regexp.Match).
// Веб-сервер должен отвечать на порт :8080.

// Пример запроса:

// curl localhost:8080/?name=John
// # hello John

// curl localhost:8080.
// # hello stranger
