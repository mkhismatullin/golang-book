package main

import "fmt"

func average1(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func average2(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total / float64(len(args))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average1(xs))
	fmt.Println(average2(xs...))
}
