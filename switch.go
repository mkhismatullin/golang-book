package main

import "fmt"

func main() {
	for i := 10; i > 0; i-- {

		var numType string

		switch i % 2 {
		case 0:
			numType = "четное"
		case 1:
			numType = "нечетное"

		}

		fmt.Println(i, numType+" число")
	}
}
