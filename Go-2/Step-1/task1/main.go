package main

import (
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(text string, user *string, status *string) []Ticket {
	result := []Ticket{}

	// Разбиваем весь текст на строки
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		// Пропускаем пустые строки
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Проверяем, что строка начинается с TICKET-
		if !strings.HasPrefix(line, "TICKET-") {
			continue
		}

		// Разбиваем строку на части
		parts := strings.Split(line, "_")
		if len(parts) != 4 {
			continue
		}

		// Парсим дату
		date, err := time.Parse("2006-01-02", parts[3])
		if err != nil {
			continue
		}

		// Создаём структуру
		ticket := Ticket{
			Ticket: parts[0],
			User:   parts[1],
			Status: parts[2],
			Date:   date,
		}

		// Фильтр по пользователю
		if user != nil && ticket.User != *user {
			continue
		}

		// Фильтр по статусу
		if status != nil && ticket.Status != *status {
			continue
		}

		// Добавляем в результат
		result = append(result, ticket)
	}

	return result
}

func main() {
	text := `TICKET-12345_Паша Попов_Готово_2024-01-01
TICKET-12346_Иван Иванов_В работе_2024-01-02
TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
TICKET-12348_Паша Попов_В работе_2024-01-04`

	user := "Паша Попов"
	tasks := GetTasks(text, &user, nil)

	for _, t := range tasks {
		println(t.Ticket, t.User, t.Status, t.Date.String())
	}
}
