package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 7
	fmt.Println("Result: ", searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	//
	// First attempt Jun 01 2024 11:04
	// Runtime 3ms, Memory 3.05MB
	//
	//start := 0
	//end := len(nums) - 1
	//
	//for start <= end {
	//	mid := (start + end) / 2
	//
	//	if nums[mid] == target {
	//		return mid
	//	}
	//
	//	if nums[mid] > target {
	//		end = mid - 1
	//	} else {
	//		start = mid + 1
	//	}
	//}
	//
	//return start
	//
	// Second attempt Jun 01 2024 11:07
	// Runtime 0ms, Memory 3.05MB
	//
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}
