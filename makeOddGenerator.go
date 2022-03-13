package main

import "fmt"

func makeOddGenerator() func() uint {
	var val uint = 1
	return func() (ret uint) {
		ret = val
		val += 2
		return
	}
}
func main() {
	odd := makeOddGenerator()
	for i := 0; i < 10; i++ {
		fmt.Println(odd())
	}
}
