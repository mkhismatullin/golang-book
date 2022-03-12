package main

import "fmt"

func main() {
	for i := 10; i > 0; i-- {
		var numType string
		if i%2 == 0 {
			numType = "четное"
		} else {
			numType = "нечетное"
		}
		fmt.Println(i, numType+" число")
	}
}
