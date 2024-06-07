package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	//
	// First attempt 07 June 2024 11:25
	// Runtime: 3ms, Memory: 3.59MB
	//
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	node := &TreeNode{Val: nums[mid]}

	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])

	return node
}
