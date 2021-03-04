package lesson2

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга.
// Площадь круга должна вводиться пользователем с клавиатуры.

func DiAndLengthOfCircle() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Circle area:\t")

	s, _ := ReadFloat64(reader)

	radius := math.Sqrt(s / math.Pi)

	di := radius * 2
	length := 2 * math.Pi * radius

	fmt.Printf("Circumference:\t%.4f\nCircle diameter : \t%.4f\n", length, di)
}
