package main

import "io"

func ReadString(r io.Reader) (string, error) {
	read, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(read), nil

}
