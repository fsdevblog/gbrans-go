package main

import (
	"fmt"
	"math"
)

// С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.

func main() {

	for {
		fmt.Printf("Введите трех значное целое число в промежутке 100...999:\t")
		var input float64

		_, _ = fmt.Scan(&input)

		if input < 100 || input > 999 {
			continue
		}
		hundreds := math.Floor(input / 100)
		tens := math.Floor((input - (hundreds * 100)) / 10)
		units := math.Floor(input - ((hundreds * 100) + (tens * 10)))

		fmt.Printf("Сотни: %d Десятки: %d Единицы: %d\n", int(hundreds), int(tens), int(units))
		fmt.Printf("More? (y/n):\t")
		var agree string
		_, _ = fmt.Scan(&agree)

		if agree == "y" || agree == "yes" {
			continue
		}
		break
	}

}
