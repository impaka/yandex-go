package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// создание файла для записи дампа горутин
	f, err := os.Create("goroutine_dump.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл:", err)
		return
	}
	defer f.Close()

	// запись дампа горутин в файл каждую секунду
	go func() {
		for {
			// получение дампа горутин
			pprof.Lookup("goroutine").WriteTo(f, 1)
			time.Sleep(time.Second)
		}
	}()

	// ваш код, в котором создаются горутины
	// ...

	// пример бесконечного цикла для демонстрации
	for {
		// здесь может быть ваша основная логика программы
		time.Sleep(1 * time.Second)
	}
}
