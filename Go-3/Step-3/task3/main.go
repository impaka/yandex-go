package main

import (
	"strconv"
	"strings"
)

func ParseHTTPStatus(statusLine string) (int, string) {
	parts := strings.Fields(statusLine)
	if len(parts) < 2 {
		return 0, ""
	}
	code, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, ""
	}
	reason := strings.Join(parts[2:], " ")
	return code, reason
}
