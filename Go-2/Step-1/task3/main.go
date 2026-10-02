package main

import "strings"

type UpperWriter struct {
	UpperString string
}

func (w *UpperWriter) Write(p []byte) (n int, err error) {
	strP := string(p)
	strP = strings.ToUpper(strP)
	w.UpperString = strP
	n = len(p)
	return n, nil
}
