package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	AnswerInvalid = errors.New("исправь свой ответ, а лучше ложись поспать")
)

func currentDayOfTheWeek() string {
	days := map[time.Weekday]string{
		time.Monday:    "Понедельник",
		time.Tuesday:   "Вторник",
		time.Wednesday: "Среда",
		time.Thursday:  "Четверг",
		time.Friday:    "Пятница",
		time.Saturday:  "Суббота",
		time.Sunday:    "Воскресенье",
	}

	now := TimeNow()
	return days[now.Weekday()]
}

func dayOrNight() string {
	now := TimeNow()
	hour := now.Hour()

	if hour >= 10 && hour <= 22 {
		return "День"
	}
	return "Ночь"
}

func nextFriday() int {
	now := TimeNow()
	today := now.Weekday()
	friday := time.Friday

	if today <= friday {
		return int(friday - today)
	}
	return int(7 - (today - friday))
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	answer = strings.ToLower(answer)
	return strings.ToLower(currentDayOfTheWeek()) == answer
}

func CheckNowDayOrNight(answer string) (bool, error) {
	answer = strings.ToLower(answer)

	if utf8.RuneCountInString(answer) != 4 {
		return false, AnswerInvalid
	}

	if answer == strings.ToLower(dayOrNight()) {
		return true, nil
	}

	return false, nil
}

func main() {
	fmt.Println(currentDayOfTheWeek())
	fmt.Println(dayOrNight())
	fmt.Println(nextFriday())
}
func TimeNow() time.Time {
	return time.Now()
}
