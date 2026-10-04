package main

import (
	"encoding/json"
)

func DeserializeStringMap(data string) (map[string]string, error) {
	var result map[string]string
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// которая принимает строку JSON и возвращает map[string]string.
// Если входная строка не является корректным JSON-объектом,
//  функция должна вернуть ошибку.
