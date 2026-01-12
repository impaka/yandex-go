package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Abs(-1.0))      // Модуль числа: 1
	fmt.Println(math.Min(-1.0, 0.1)) // Минимальное число: -1
	fmt.Println(math.Max(-1.0, 0.1)) // Максимальное число: 0.1
	fmt.Println(math.Sqrt(2.0))      // Корень: 1.4142135623730951
	fmt.Println(math.Sin(45.0))      // Синус: 0.8509035245341183
	fmt.Println(math.Pow(2.0, 3.0))  // Степень: 8
	fmt.Println(math.Ceil(5.3))      // Наименьшее целое, большее или равное числу: 6
	fmt.Println(math.Floor(5.3))     // Наибольшее целое, меньшее или равное числу: 5
	fmt.Println(math.Round(5.5))     // Ближайшее целое число, округляя половину от нуля: 6
}
