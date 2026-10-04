package main

import "time"

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	defer close(prime_nums)
	time.AfterFunc(100*time.Millisecond, func() { close(stop) })
	for i := 2; i <= N; i++ {
		isPrime := true
		for d := 2; d*d <= i; d++ {
			if i%d == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			select {
			case prime_nums <- i:

			case <-stop:
				return
			}
		}
	}
}
