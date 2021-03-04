package lesson2

import (
	"fmt"
	"math"
)

// С клавиатуры вводится трехзначное число.
// Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.

func DigitsToWords() {
	fmt.Printf("Enter a three-digit integer in the range 100...999:\t")

	var input float64

	_, _ = fmt.Scan(&input)

	if input < 100 || input > 999 {
		DigitsToWords()
		return
	}

	hundreds := math.Floor(input / 100)
	tens := math.Floor((input - (hundreds * 100)) / 10)
	units := math.Floor(input - ((hundreds * 100) + (tens * 10)))

	fmt.Printf("Hunderds: %d Tens: %d Units: %d\n", int(hundreds), int(tens), int(units))
}
