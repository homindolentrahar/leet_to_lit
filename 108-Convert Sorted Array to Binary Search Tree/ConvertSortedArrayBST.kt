fun main() {
    val nums = intArrayOf(-10, -3, 0, 5, 9)

    ConvertSortedArrayBST.sortedArrayToBST(nums)
}

class TreeNode(var `val`: Int) {
    var left: TreeNode? = null
    var right: TreeNode? = null
}

object ConvertSortedArrayBST {
    //
    // First attempt 07 June 2024
    // Runtime: 169ms, 37.66MB
    //
    fun sortedArrayToBST(nums: IntArray): TreeNode? {
        if (nums.isEmpty()) {
            return null
        }

        val mid = nums.size / 2
        val node = TreeNode(nums[mid])

        node.left = sortedArrayToBST(nums.sliceArray(0 until mid))
        node.right = sortedArrayToBST(nums.sliceArray(mid + 1 until nums.size))

        return node
    }
}