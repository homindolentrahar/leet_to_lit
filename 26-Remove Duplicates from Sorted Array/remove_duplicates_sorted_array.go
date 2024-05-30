package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Println("Result: ", removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	// First attempt May 30, 2024 20:32
	// Runtime 8ms, Memory 4.47MB
	//
	unique := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[unique] = nums[i]
			unique++
		}
	}

	return unique
}
