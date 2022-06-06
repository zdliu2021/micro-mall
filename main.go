package main

import (
	"fmt"
)

func test(params ...int) {
	fmt.Println(params)
}

func main() {
	s := []int{1, 2, 3}
	test(s...)
}
