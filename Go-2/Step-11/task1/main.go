package main

func Sum[T int | float64 | string](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}
