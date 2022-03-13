package main

import "fmt"

func calcFibbonachi() func() uint64 {
	var prev2, prev1 uint64 = 0, 0
	var n = 0
	return func() (result uint64) {
		result = prev2 + prev1
		if n == 0 {
			prev2 = 0
		} else if n == 1 {
			prev1 = 1
		} else {
			prev2 = prev1
			prev1 = result
		}
		n += 1
		//fmt.Println("F#", n, " : prev1 = ", prev1, " : prev2 = ", prev2)
		return
	}
}
func main() {
	fib := calcFibbonachi()
	var count int
	fmt.Print("Укажите количество чисел Фиббоначчи для расчета (<95):")
	fmt.Scanf("%d", &count)
	if count >= 95 {
		count = 94
	}
	for i := 0; i < count; i++ {
		fmt.Println("F#", i, " = ", fib())
	}
}
