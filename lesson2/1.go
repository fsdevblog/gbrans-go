package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Напишите программу для вычисления площади прямоугольника. Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Длина:\t")
		length, _ := readInput(reader)

		fmt.Printf("Ширина:\t")
		width, _ := readInput(reader)

		fmt.Printf("Площадь:\t%.2f\n", length*width)
		fmt.Printf("More? (y/n):\t")

		var agree string
		_, _ = fmt.Scan(&agree)

		if agree == "y" || agree == "yes" {
			continue
		}

		break
	}

}

func readInput(reader *bufio.Reader) (float64, error) {
	str, _ := reader.ReadString('\n')
	str = strings.ReplaceAll(str, ",", ".")
	str = strings.ReplaceAll(str, "\n", "")
	return strconv.ParseFloat(str, 32)
}
