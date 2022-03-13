package main

import "fmt"

func getMin(args ...int) int {
	var min int
	if len(args) > 0 {
		min = args[0]
		for _, v := range args {
			if min > v {
				min = v
			}
		}
	}
	return min
}
func main() {
	var mySlice []int
	mySlice = []int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	fmt.Println(getMin(mySlice...))
}
