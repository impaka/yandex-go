package main

import "sort"

func SortAndMerge(left, right []int) []int {
	// сортируем входные срезы (изменяет их)
	sort.Ints(left)
	sort.Ints(right)

	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// добавить оставшиеся элементы
	if i < len(left) {
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}

func main() {

}
