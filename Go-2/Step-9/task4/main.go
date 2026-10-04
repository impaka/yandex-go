package main

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	if len(str1) != len(str2) {
		return false
	}

	// Convert strings to slices of runes for sorting
	r1 := []rune(str1)
	r2 := []rune(str2)

	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})

	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	return string(r1) == string(r2)
}

// func main() {
// 	res := AreAnagrams("кот", "ток")
// 	fmt.Println(res)
// 	res = AreAnagrams("папа", "мама")
// 	fmt.Println(res)
// }
