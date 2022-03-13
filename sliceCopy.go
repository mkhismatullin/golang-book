package main

import "fmt"

func main() {
	var slice1 = []string{"1", "2", "3", "4", "5", "6", "7"}
	slice2 := slice1[2:4]
	for _, value := range slice2 {
		fmt.Println(value)
	}
	fmt.Println(slice1, slice2)
}
