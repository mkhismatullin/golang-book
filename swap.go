package main

import "fmt"

func swap(xPtr, yPtr *int) {
	/*temp := new(int)
	*temp = *xPtr
	*xPtr = *yPtr
	*yPtr = *temp*/
	*xPtr, *yPtr = *yPtr, *xPtr
}
func main() {
	var x, y = 1, 2
	fmt.Println("before:", x, ",", y)
	swap(&x, &y)
	fmt.Println("after :", x, ",", y)
}
