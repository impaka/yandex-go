package main

import (
	"io"
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Seek(int64(pos), io.SeekStart)
	if err != nil {
		return
	}
	file.Write([]byte(val))
}
