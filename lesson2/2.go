package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Площадь круга:\t")
		s, _ := readInput2(reader)

		radius := math.Sqrt(s / math.Pi)

		di := radius * 2
		length := 2 * math.Pi * radius

		fmt.Printf("Длина окружности:\t%.4f\nДиаметр окружности: \t%.4f %f\n\n", length, di, radius)
		fmt.Printf("More? (y/n):\t")

		var agree string
		_, _ = fmt.Scan(&agree)

		if agree == "y" || agree == "yes" {
			continue
		}

		break
	}

}

/*
Не разбирался еще с package, IDE ругается на дублирование функции из 1.go
поэтому дал название readInput2
*/

func readInput2(reader *bufio.Reader) (float64, error) {
	str, _ := reader.ReadString('\n')
	str = strings.ReplaceAll(str, ",", ".")
	str = strings.ReplaceAll(str, "\n", "")
	return strconv.ParseFloat(str, 32)
}
