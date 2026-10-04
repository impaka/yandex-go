package main

import "strings"

func MakeCurlCommand(method, url, headers, body string) string {
	var b strings.Builder
	b.WriteString("curl ")
	if method != "GET" {
		b.WriteString("-X ")
		b.WriteString(method)
		b.WriteString(" ")
	}
	for _, h := range strings.Split(headers, "\n") {
		if h == "" {
			continue
		}
		b.WriteString("-H '")
		b.WriteString(h)
		b.WriteString("' ")
	}

	if body != "" {
		b.WriteString("--data '")
		b.WriteString(body)
		b.WriteString("' ")
	}
	b.WriteString(url)
	return b.String()

}

// Реализуйте функцию MakeCurlCommand(method, url, headers, body string) string,
// которая формирует строку команды для утилиты curl на основе переданных параметров:

// method — HTTP-метод (например, "GET", "POST"). Для "GET" параметр -X можно не указывать
// url — полный адрес (например, "https://example.com/api/data")
// headers — дополнительные заголовки (каждый заголовок заканчивается символом новой строки, например,
//  "Content-Type: application/json\nAuthorization: Bearer xyz\n")
// body — тело запроса (использовать только если оно не пустое; передавать через --data '...')
// Функция должна вернуть строку команды, которую можно выполнить в терминале.
// Каждый заголовок должен указываться через -H '...', тело через --data '...',
// URL — последний параметр. Заголовки и тело должны корректно экранироваться одинарными кавычками.
