package main

import "fmt"

func main() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(list[2:5])
	//[2 3 4]

	slice := make([]int, 3, 9)
	fmt.Println(len(slice))
	//3

	list = []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := list[0]
	for _, value := range list {
		if value < min {
			min = value
		}
	}
	fmt.Println(min)
	//9
}
