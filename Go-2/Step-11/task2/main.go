package main

func Filter[T any](arr []T, predicate func(T) bool) []T {
	var total []T
	for _, v := range arr {
		if predicate(v) {
			total = append(total, v)
		}
	}
	return total

}
