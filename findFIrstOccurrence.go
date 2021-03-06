package main

// Find the first occurrence in sorted array
func FindFirstOccurrence(arr []int, target int) int {

	low := 0
	high := len(arr) - 1
	result := -1

	for low <= high {
		mid := (low + high) >> 1
		if target == arr[mid] {
			result = mid
			high = mid - 1
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}
