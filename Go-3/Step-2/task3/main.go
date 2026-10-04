package main

import (
	"encoding/json"
	"io"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func DecodeStudentFromReader(r io.Reader) (Student, error) {
	var student Student
	input, err := io.ReadAll(r)
	if err != nil {
		return Student{}, err
	}
	err = json.Unmarshal(input, &student)
	if err != nil {
		return Student{}, err
	}
	return student, nil
}
