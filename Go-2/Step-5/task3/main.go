package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"
)

// Напишите функцию

// func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse
// которая одновременно получает данные из переданных urls методом GET.

// Используйте контекст ctx, чтобы ограничить время запроса и отменить его
//  при ожидания свыше timeout.
//  В случае ошибки верните её в соответствующем объекте APIResponse.
// При превышении времени ожидания должна быть ошибка context.DeadlineExceeded.

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func fetchOne(ctx context.Context, url string) *APIResponse {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &APIResponse{URL: url, Err: err}
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return &APIResponse{URL: url, Err: context.DeadlineExceeded}
		}
		return &APIResponse{URL: url, Err: err}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &APIResponse{URL: url, Err: err}
	}
	return &APIResponse{
		URL:        url,
		Data:       string(body),
		StatusCode: resp.StatusCode,
		Err:        nil,
	}

}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	newCTX, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	result := make([]*APIResponse, len(urls))
	wg := sync.WaitGroup{}
	for i, url := range urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			result[i] = fetchOne(newCTX, url)
		}(i, url)

	}
	wg.Wait()
	return result
}
