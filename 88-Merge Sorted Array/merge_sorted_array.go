package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3

	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	//
	// Failed attempt 06 June 2024 13:56
	// 3/59 testcases passed
	//
	//lastIndex := m + n - 1
	//lastM := m - 1
	//lastN := n - 1
	//
	//if lastM < 0 {
	//	nums1 = nums2
	//	fmt.Println("Nums: ", nums1)
	//	return
	//}
	//
	//for i := lastIndex; i > 0; i-- {
	//	if nums2[lastN] > nums1[lastM] {
	//		nums1[lastIndex] = nums2[lastN]
	//		lastIndex--
	//		lastN--
	//	} else {
	//		nums1[lastIndex] = nums1[lastM]
	//		lastM--
	//		lastIndex--
	//	}
	//}
	//
	//fmt.Println("Nums: ", nums1)
	//
	// Second attempt 06 June 2024 14:22
	// Runtime: 2ms, Memory: 2.40MB
	//
	lastPointer := m + n - 1

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[lastPointer] = nums1[m-1]
			m--
		} else {
			nums1[lastPointer] = nums2[n-1]
			n--
		}

		lastPointer--
	}

	for n > 0 {
		nums1[lastPointer] = nums2[n-1]
		n, lastPointer = n-1, lastPointer-1
	}

	fmt.Println("Nums: ", nums1)
	//
	// Third attempt 06 June 2024 14:27
	// Runtime: 2ms, Memory: 2.44MB
	//
	//lastPointer := m + n - 1
	//pointerM := m - 1
	//pointerN := n - 1
	//
	//for pointerM >= 0 && pointerN >= 0 {
	//	if nums1[pointerM] > nums2[pointerN] {
	//		nums1[lastPointer] = nums1[pointerM]
	//		pointerM--
	//	} else {
	//		nums1[lastPointer] = nums2[pointerN]
	//		pointerN--
	//	}
	//
	//	lastPointer--
	//}
	//
	//for pointerN >= 0 {
	//	nums1[lastPointer] = nums2[pointerN]
	//	pointerN--
	//	lastPointer--
	//}
	//
	//fmt.Println("Nums: ", nums1)
}
