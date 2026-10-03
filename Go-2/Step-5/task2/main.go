package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data       string // тело ответа
	StatusCode int    // код ответа
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(newCtx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, context.DeadlineExceeded
		}
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	apiResponce := APIResponse{
		Data:       string(body),
		StatusCode: resp.StatusCode,
	}
	return &apiResponce, nil
}
