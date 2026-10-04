package main

import "strings"

func BuildHTTPRequest(method, url, host, headers, body string) string {
	var b strings.Builder
	b.WriteString(method)
	b.WriteString(" ")
	b.WriteString(url)
	b.WriteString(" HTTP/1.1\r\n")

	b.WriteString("Host: ")
	b.WriteString(host)
	b.WriteString("\r\n")

	b.WriteString(headers)
	b.WriteString("\r\n")
	b.WriteString(body)
	return b.String()

}
