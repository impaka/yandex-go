package main

import (
	"context"
	"sync"
)

type pair struct {
	idx   int
	value int
}

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error) {
	wg := sync.WaitGroup{}

	if err := ctx.Err(); err != nil {
		return nil, ctx.Err()
	}
	jobs := make(chan int)
	go func() {
		defer close(jobs)
		for i := range inputs {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}
	}()
	result := make(chan pair, len(inputs))
	go func() {
		wg.Wait()
		close(result)
	}()
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case idx, ok := <-jobs:
					if !ok {
						return
					}
					result <- pair{idx: idx, value: fn(inputs[idx])}
				}
			}
		}()
	}

	wg.Wait()

	out := make([]int, len(inputs))
	for r := range result {
		out[r.idx] = r.value
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return out, nil
}
