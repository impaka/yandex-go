package main

import (
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {

	fileOpen, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer fileOpen.Close()
	_, err = fileOpen.Seek(int64(startpos), io.SeekStart)
	if err != nil {
		return err
	}
	var resultReadFile string
	buffer := make([]byte, 100)
	for {
		n, err := fileOpen.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		resultReadFile += string(buffer[:n])

	}
	fileCreate, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	defer fileCreate.Close()

	fileCreate.WriteString(resultReadFile)
	return nil
}
