package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var queueMas = make([]string, 5)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.Fields(line)

		// Команды
		if len(parts) == 1 {
			cmd := parts[0]
			switch cmd {
			case "очередь":
				printQueue()
			case "количество":
				quantity()
			case "конец":
				printQueue()
				return
			default:
				fmt.Println("Некорректный ввод")
			}
			continue
		}

		// Запись: имя + номер
		if len(parts) == 2 {
			name := parts[0]
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Запись на место номер %s невозможна: некорректный ввод\n", parts[1])
				continue
			}
			record(num, name)
			continue
		}

		// Всё остальное — некорректный ввод
		fmt.Println("Некорректный ввод")
	}
}

func printQueue() {
	for i, name := range queueMas {
		if name == "" {
			fmt.Printf("%d. -\n", i+1)
		} else {
			fmt.Printf("%d. %s\n", i+1, name)
		}
	}
}

func quantity() {
	occupied := 0
	for _, v := range queueMas {
		if v != "" {
			occupied++
		}
	}
	free := 5 - occupied
	fmt.Printf("Осталось свободных мест: %d\nВсего человек в очереди: %d\n", free, occupied)
}

func record(num int, name string) {
	if num < 1 || num > 5 || name == "" {
		fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", num)
		return
	}

	// Проверка переполнения
	full := true
	for _, v := range queueMas {
		if v == "" {
			full = false
			break
		}
	}
	if full {
		fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", num)
		return
	}

	// Проверка занятости места
	if queueMas[num-1] != "" {
		fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", num)
		return
	}

	queueMas[num-1] = name
}
