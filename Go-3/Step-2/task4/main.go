package main

import (
	"encoding/json"
	"io"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func EncodeStudentsToWriter(w io.Writer, students []Student) error {
	studentsJson, err := json.Marshal(students)
	if err != nil {
		return err
	}
	_,err = w.Write(studentsJson) 
	if err != nil {
		return err
	}
	return nil
}
