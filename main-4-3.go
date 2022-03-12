package main

import "fmt"

func main() {
	fmt.Print("Температура по Фаренгейту:")
	var tempF float64
	fmt.Scanf("%f", &tempF)
	tempC := (tempF - 32) * 5 / 9
	fmt.Println("Температура по Цельсию:", tempC)
}
