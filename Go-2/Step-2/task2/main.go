package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	var current int
	file, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if current == lineNum {
			return scanner.Text()
		}
		current++

	}
	return ""
}
