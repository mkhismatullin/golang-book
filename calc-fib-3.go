package main

import "fmt"

func calcFib(n int) func() uint {
	var prev2, prev1 uint //= 0, 0
	return func() (result uint) {
		//result = prev2 + prev1
		fmt.Println("F#", n, " : prev1 = ", prev1, " : prev2 = ", prev2)
		if n == 0 {
			prev2 = 0
			result = 0
		} else if n == 1 {
			prev1 = 1
			result = 1
		} else {
			result = prev2 + prev1
			prev2 = prev1
			prev1 = result
		}
		//fmt.Println("F#", n, " : prev1 = ", prev1, " : prev2 = ", prev2)
		return
	}
}
func main() {
	//var fib func() uint
	for i := 0; i < 10; i++ {
		//fib = calcFib(i)
		fmt.Println(calcFib(i)())
	}
}
