package lesson3

import (
	"fmt"
	"math/big"
)

// Написать приложение, которое ищет все простые числа от 0 до N включительно.
// Число N должно быть задано из стандартного потока ввода. P.S. в math есть подсказка как это сделать
func Prime() {

	var n int64
	_, _ = fmt.Scan(&n)

	var primes []int64

	for i := int64(0); i <= n; i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			primes = append(primes, i)
		}
	}
	fmt.Printf("Primes: %v\n", primes)
}
