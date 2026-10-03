package main

import (
	"strings"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	var count int
	for i := range questions {
		select {
		case ans := <-answerCh:
			if strings.EqualFold(strings.TrimSpace(ans), strings.TrimSpace(answers[i])) {
				count++
			}
		case <-time.After(time.Second):
			continue
		}

	}
	return count

}
