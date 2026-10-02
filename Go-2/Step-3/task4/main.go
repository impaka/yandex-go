package main

func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for _, val := range nums {
		ch <- val
	}
	return ch
}
