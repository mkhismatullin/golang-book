package main

import "fmt"

func half(input int) (float64, bool) {
	output := float64(input) / 2
	evenFlag := input%2 == 0
	return output, evenFlag
}

func main() {
	var input int
	fmt.Print("Введите число:")
	fmt.Scanf("%d", &input)
	output, evenFlag := half(input)
	fmt.Println("Половина = ", output, func() string {
		if evenFlag {
			return "Четное"
		} else {
			return "Нечетное"
		}
	}())
}
