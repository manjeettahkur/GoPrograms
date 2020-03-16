package main

import "fmt"

func main() {
	//arr := []int{8, 9, 10, 1, 2, 3}
	//l := len(arr) - 1
	//res := FindRotationalCount(arr, l)

	//arr2 := []int{3, 4, 5, 1, 2}
	//res := FindMin(arr2)
	//fmt.Println(res)

	arr3 := []int{2, 5, 5, 5, 6, 6, 8, 9, 9, 9}
	res := FindFirstOccurrence(arr3, 8)
	res1 := FindLastOccurrence(arr3, 8)
	fmt.Println(res, "::: ", res1)
}
