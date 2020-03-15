package main

import (
	"fmt"
)

func main() {
	arr := []int{8, 9, 10, 1, 2, 3}
	l := len(arr) - 1
	res := FindRotationalCount(arr, l)
	fmt.Println(res)
}
