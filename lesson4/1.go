package lesson4

func InjectionSort(ary []int) {
	aryLength := len(ary)
	if aryLength < 2 {
		return
	}
	// Цикл всего массива
	for i := 1; i < aryLength; i++ {
		// цикл подмассива с пройдеными элементами
		for j := i; j > 0 && ary[j-1] > ary[j]; j-- {
			ary[j-1], ary[j] = ary[j], ary[j-1]
		}
	}
}
