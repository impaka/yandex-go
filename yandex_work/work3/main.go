package main

import (
	"fmt"
	"time"
)

func main() {
	var timeNow string
	var lastName string
	var firstName string
	var patronymicName string
	var oneMoney float64
	var twoMoney float64
	var threeMoney float64

	fmt.Scanln(&timeNow)
	fmt.Scanln(&lastName)
	fmt.Scanln(&firstName)
	fmt.Scanln(&patronymicName)
	fmt.Scanln(&oneMoney)
	fmt.Scanln(&twoMoney)
	fmt.Scanln(&threeMoney)

	// Парсим дату
	resultTime, err := time.Parse("02.01.2006", timeNow)
	if err != nil {
		fmt.Println("Ошибка парсинга даты:", err)
		return
	}

	// Прибавляем 15 дней
	resultTime = resultTime.AddDate(0, 0, 15)

	// Формируем дату в нужном формате
	day := resultTime.Day()
	month := int(resultTime.Month())
	year := resultTime.Year()
	formattedDate := fmt.Sprintf("%02d.%02d.%d", day, month, year)

	// Считаем сумму
	total := oneMoney + twoMoney + threeMoney
	kopecks := int(total * 100)
	rubles := kopecks / 100
	kopecks = kopecks % 100
	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n"+"Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\n"+"Общая сумма выплат составит %d руб. %d коп.\n\n"+"С уважением,\n"+"Гл. бух. Иванов А.Е.", firstName,
		lastName,
		patronymicName,
		formattedDate,
		rubles,
		kopecks,
	)
}

// Уважаемый, Иванов Андрей Валерьевич, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.
// Дата подписания договора: 25.04.2005. Просим вас подойти в офис в любое удобное для вас время в этот день.
// Общая сумма выплат составит 35122 руб. 0 коп.

// С уважением,
// Гл. бух. Иванов А.Е.
// Ввод
// 10.04.2005
// Андрей
// Иванов
// Валерьевич
// 15000
// 19999
// 123
