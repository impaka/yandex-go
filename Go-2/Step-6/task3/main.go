package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type student struct {
	name  string
	grade int
}

func BestStudents(names []string) (string, error) {
	wg := sync.WaitGroup{}
	var mx sync.Mutex
	var firstErr error
	var students []student

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
			students = append(students, student{n, grRes})
			mx.Unlock()
		}(name)
	}
	wg.Wait()
	if firstErr != nil {
		return "", firstErr
	}
	var sum int
	for _, student := range students {
		sum += student.grade
	}
	var best []string
	average := sum / len(students)
	for _, student := range students {
		if student.grade > average {
			best = append(best, student.name)

		}
	}
	sort.Strings(best)
	return strings.Join(best, ","), nil
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
