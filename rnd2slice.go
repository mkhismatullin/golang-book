package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	slice := make([]string, 10)
	curTime := time.Now()
	// Each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	randFromS, _ := strconv.Atoi(curTime.Format("5"))
	rand.Seed(int64(randFromS))
	for i := 0; i < len(slice); i++ {
		slice[i] = string(rand.Intn(len(slice)) + 48)
	}
	fmt.Println(slice)
}
