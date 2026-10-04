package main

import "strings"

func BuildHTTPResponse(statusLine, headers, body string) string {
	var b strings.Builder
	b.WriteString(statusLine)
	b.WriteString("\r\n")
	b.WriteString(headers)
	b.WriteString("\r\n")
	b.WriteString(body)

	return b.String()
}
