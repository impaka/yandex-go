package main

import (
	"slices"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	numbers := []int{3, 1, 2}
	SortIntegers(numbers)
	if !slices.Equal(numbers, []int{1, 2, 3}) {
		t.Errorf("SortIntegers = %v, ожидалось %v", numbers, []int{1, 2, 3})
	}
}
