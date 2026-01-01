package main

import (
	"fmt"
	"sort"
	"strings"
)

func AnalyzeText(text string) {
	sliceText, lenText := splitText(text)
	uniqueMap := unique(sliceText)
	lenUnique := len(uniqueMap)
	frequentWord, maxCount := frequentWord(uniqueMap)
	// Получаем топ‑5
	top5 := getTopWords(uniqueMap, 5)

	// Красивый вывод
	fmt.Printf("Количество слов: %d\n", lenText)
	fmt.Printf("Количество уникальных слов: %d\n", lenUnique)
	fmt.Printf("Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n", frequentWord, maxCount)
	fmt.Println("Топ-5 самых часто встречающихся слов:")

	for _, w := range top5 {
		fmt.Printf("\"%s\": %d раз\n", w, uniqueMap[w])
	}
}
func frequentWord(uniqueMap map[string]int) (string, int) {
	// Находим самое частое слово
	var frequentWord string
	var maxCount int
	first := true
	for k, v := range uniqueMap {
		if first || v > maxCount {
			frequentWord = k
			maxCount = v
			first = false
		}
	}
	return frequentWord, maxCount
}
func unique(sliceText []string) map[string]int {
	uniqueMap := make(map[string]int)
	for _, word := range sliceText {
		uniqueMap[word]++
	}
	return uniqueMap
}
func splitText(text string) ([]string, int) {
	text = cleanText(text)
	sliceText := strings.Fields(text)
	lenText := len(sliceText)
	return sliceText, lenText
}
func cleanText(text string) string {
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, "!", " ")
	text = strings.ReplaceAll(text, "?", " ")
	text = strings.ToLower(text)
	return text
}
func getTopWords(wordMap map[string]int, n int) []string {
	type pair struct {
		word  string
		count int
	}

	// Собираем пары (слово + количество) из мапы
	pairs := make([]pair, 0, len(wordMap))
	for w, c := range wordMap {
		pairs = append(pairs, pair{word: w, count: c})
	}

	// Сортируем по count по убыванию
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	// Берём первые n слов
	if n > len(pairs) {
		n = len(pairs)
	}
	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, pairs[i].word)
	}

	return res
}

func main() {
	testText := "Привет, мир! Как дела? дела"
	AnalyzeText(testText)
}
