package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

func CompareList(names []string) (map[string]string, error) {
	wg := sync.WaitGroup{}
	var mx sync.Mutex
	var firstErr error
	grades := make(map[string]int, len(names))

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
			grades[n] = grRes
			mx.Unlock()
		}(name)
	}
	wg.Wait()
	if firstErr != nil {
		return nil, firstErr
	} else if len(grades) == 0 {
		return nil, nil
	}
	var sum int
	for _, grade := range grades {
		sum += grade
	}
	average := sum / len(grades)
	result := make(map[string]string, len(grades))
	for name, g := range grades {
		switch {
		case g > average:
			result[name] = ">"
		case g < average:
			result[name] = "<"
		case g == average:
			result[name] = "="
		}
	}
	return result, nil
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
