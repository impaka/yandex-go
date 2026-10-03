package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func Average(names []string) (int, error) {
	wg := sync.WaitGroup{}
	var grades []int
	var mx sync.Mutex
	var firstErr error
	for _, name := range names {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			res, err := request(n)
			if err != nil {
				mx.Lock()
				if firstErr == nil {
					firstErr = err
				}
				mx.Unlock()
				return
			}

			grRes, err := strconv.Atoi(res)
			if err != nil {
				mx.Lock()
				if firstErr == nil {
					firstErr = err
				}
				mx.Unlock()
				return
			}
			mx.Lock()
			grades = append(grades, grRes)
			mx.Unlock()
		}(name)
	}
	wg.Wait()
	if firstErr != nil {
		return 0, firstErr
	}
	if len(grades) == 0 {
		return 0, nil
	}
	var sum int
	for _, num := range grades {
		sum += num
	}
	return sum / len(grades), nil
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
