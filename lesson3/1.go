package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operand = string

const (
	PLUS     = Operand("+")
	MINUS    = Operand("-")
	DIVIDE   = Operand("/")
	INCREASE = Operand("*")
	POW      = Operand("**")
)

type Expression struct {
	firstOperator  float64
	operand        Operand
	secondOperator float64
}

func (e *Expression) toString() string {
	return fmt.Sprintf("%f %s %f", e.firstOperator, e.operand, e.secondOperator)
}

func (e *Expression) calc() (float64, error) {
	switch e.operand {
	case PLUS:
		return e.firstOperator + e.secondOperator, nil
	case MINUS:
		return e.firstOperator - e.secondOperator, nil
	case INCREASE:
		return e.firstOperator * e.secondOperator, nil
	case DIVIDE:
		return e.firstOperator / e.secondOperator, nil
	case POW:
		return math.Pow(e.firstOperator, e.secondOperator), nil
	default:
		return 0, fmt.Errorf("unsupported operand %v", e.operand)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please input expression:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		normalizeInput(&input)
		result, err := calculate(input)

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("RESULT: %s\n", result)
		fmt.Printf("More? (y/n):\t")
		var agree string
		_, _ = fmt.Scan(&agree)

		if agree == "y" || agree == "yes" {
			continue
		}
		break
	}
}

func calculate(str string) (string, error) {
	parts := strings.Fields(str)

	// Когда осталась одна часть выражения, рекурсия закончена, результат найден
	if len(parts) == 1 {
		return parts[0], nil
	}

	// В выражении должно быть 3 части (оператор, операнд, операнд)
	if len(parts) < 3 {
		return "", errors.New("something wrong at parse expression step")
	}

	// Далее 2 однотипные операции, я не понимаю как это написать DRY

	firstOp, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return "", err
	}
	secondOp, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return "", err
	}
	expression := Expression{
		operand:        parts[1],
		firstOperator:  firstOp,
		secondOperator: secondOp,
	}
	result, err := expression.calc()
	if err != nil {
		return "", err
	}
	newString := strings.Replace(str, parts[0]+" "+parts[1]+" "+parts[2], fmt.Sprintf("%f", result), 1)
	return calculate(newString)
}

// Чистим ввод от левых символов и подготавливаем выражение к парсингу
func normalizeInput(str *string) {
	clean := regexp.MustCompile(`[^\d\+\-/\*\.]`).ReplaceAll([]byte(*str), []byte(""))
	clean = regexp.MustCompile(`([\+\-/\*])`).ReplaceAll(clean, []byte(" $1 "))
	res := regexp.MustCompile(`\s\s+`).ReplaceAll(clean, []byte(" "))
	*str = strings.ReplaceAll(string(res), "* *", "**")
}