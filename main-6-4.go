package main

import "fmt"

func main() {
	var slice1 = []string{"1", "2", "3", "4", "5", "6", "7"}
	slice1 = append(slice1[:3], slice1[4:]...)
	fmt.Println(slice1)
}
