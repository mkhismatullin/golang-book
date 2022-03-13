package main

import "fmt"

func calcFibbonachi(n int) int {
	var result int
	if n <= 1 {
		result = n
	} else {
		result = calcFibbonachi(n-1) + calcFibbonachi(n-2)
	}
	return result
}
func main() {
	var count int
	fmt.Print("Укажите количество чисел Фиббоначчи для расчета: ")
	fmt.Scanf("%d", &count)
	for i := 0; i < count; i++ {
		fmt.Println("F#", i, " = ", calcFibbonachi(i))
	}
}
