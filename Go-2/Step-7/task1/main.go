package main

import (
	"errors"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	if n < 0 {
		return 0, errors.New("n must be non-negative")

	}
	ch := make(chan int, 1)
	go func() {
		ch <- fibonacci(n)
	}()

	select {
	case res := <-ch:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}

}

func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
