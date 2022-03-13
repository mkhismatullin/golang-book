package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		errMsg := recover()
		fmt.Println(errMsg)
	}()
	panic(time.Now())
}
