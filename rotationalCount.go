package main

import "fmt"

func FindRotationalCount(arr []int, l int) int {

	low := 0
	high := l

	for low <= high {
		if arr[low] <= arr[high] {
			return low
		}

		mid := low + high/2
		next := (mid + 1) % l
		prev := (mid - 1 + l) % l

		fmt.Println("prev", arr[prev], "mid", arr[mid], "next", arr[next])

		if arr[mid] <= arr[next] && arr[mid] <= arr[prev] {
			return mid
		} else if arr[mid] <= arr[high] {
			high = mid - 1
		} else if arr[mid] >= arr[low] {
			low = mid + 1
		}
	}
	return -1
}
