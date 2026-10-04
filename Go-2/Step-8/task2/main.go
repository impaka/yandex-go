package main

import (
	"context"
	"io"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	defer close(result)
	fileOpen, err := os.Open(path)
	if err != nil {
		return
	}
	defer fileOpen.Close()
	var fileBuf []byte
	buf := make([]byte, 100)
	for {
		if ctx.Err() != nil {
			return
		}

		n, err := fileOpen.Read(buf)
		if n > 0 {
			fileBuf = append(fileBuf, buf[:n]...)
		}
		if err == io.EOF {
			select {
			case result <- fileBuf:
			case <-ctx.Done():
			}
			return
		}
		if err != nil {
			return
		}

	}

}
