package main

import (
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func GetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	timeout time.Duration,
) error {
	newCTX, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if err := newCTX.Err(); err != nil {
		return newCTX.Err()
	}
	result := []Ticket{}
	text, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	if err := newCTX.Err(); err != nil {
		return err
	}
	// Разбиваем весь текст на строки
	lines := strings.Split(string(text), "\n")
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
		switch parts[2] {
		case "Готово", "В работе", "Не будет сделано":
		default:
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

	resultJson, err := json.Marshal(result)
	if err != nil {
		return err
	}
	_, err = w.Write(resultJson)
	if err != nil {
		return err
	}

	return nil

}
