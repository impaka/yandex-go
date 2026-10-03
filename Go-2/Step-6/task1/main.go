package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func Compare(name1, name2 string) (string, error) {
	wg := sync.WaitGroup{}
	var grade1, grade2 string
	var err1, err2 error

	wg.Add(1)
	go func() {
		defer wg.Done()
		grade1, err1 = request(name1)

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		grade2, err2 = request(name2)

	}()
	wg.Wait()
	if err1 != nil {
		return "", err1
	}
	if err2 != nil {
		return "", err2
	}
	g1, err := strconv.Atoi(grade1)
	if err != nil {
		return "", err
	}
	g2, err := strconv.Atoi(grade2)
	if err != nil {
		return "", err
	}
	switch {
	case g1 > g2:
		return ">", nil
	case g1 < g2:
		return "<", nil
	default:
		return "=", nil

	}
}

func request(name string) (string, error) {
	q := "http://localhost:8082/mark?" + url.Values{"name": {name}}.Encode()
	req, err := http.NewRequest(http.MethodGet, q, nil)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil

}
