package lesson2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RectangleArea() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Length:\t")

	length, _ := ReadFloat64(reader)

	fmt.Printf("Width:\t")

	width, _ := ReadFloat64(reader)

	fmt.Printf("Area:\t%.2f\n", length*width)
}

func ReadFloat64(reader *bufio.Reader) (float64, error) {
	str, _ := reader.ReadString('\n')
	str = strings.ReplaceAll(str, ",", ".")
	str = strings.ReplaceAll(str, "\n", "")

	return strconv.ParseFloat(str, 32)
}
