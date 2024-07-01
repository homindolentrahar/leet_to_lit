fun main() {
    val nums = intArrayOf(3, 2, 3)

    println("Output: ${MajorityElement.majorityElement(nums)}")
}

object MajorityElement {
    fun majorityElement(nums: IntArray): Int {
        //
        // First attempt 01 July 2024 16:54
        // Runtime: 245ms, Memory: 41.80MB
        // Time complexity: O(n), Space complexity: O(1)
        // Boyer-Moore Voting Algorithm
        //
        var candidate = 0
        var count = 0

        for (num in nums) {
            if (count == 0) {
                candidate = num
            }

            if (candidate == num) {
                count++
            } else {
                count--
            }
        }

        return candidate
    }
}