package main

import (
	"fmt"
)

func main() {
	var oneInput string
	var twoInput int

	fmt.Scan(&oneInput, &twoInput)

	if oneInput == "-" {
		twoInput = -twoInput
	}
	if twoInput > 20 {
		fmt.Println("Стоит надеть майку и шорты")
	} else if twoInput >= 10 && twoInput <= 20 {
		fmt.Println("Стоит надеть штаны и кофту")
	} else if twoInput >= -5 && twoInput <= 9 || twoInput == 0 {
		fmt.Println("Стоит надеть куртку")
	} else if twoInput <= -6 {
		fmt.Println("Стоит надеть зимнюю куртку")
	}
}
