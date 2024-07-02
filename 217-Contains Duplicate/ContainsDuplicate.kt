fun main() {
    val nums = intArrayOf(1, 2, 3, 1)

    println("Output: ${ContainsDuplicate.containsDuplicate(nums)}")
}

object ContainsDuplicate {
    fun containsDuplicate(nums: IntArray): Boolean {
        //
        // First attempt 02 July 2024 16:54
        // Runtime: 245ms, Memory: 41.80MB
        // Time complexity: O(n), Space complexity: O(1)
        // HashSet
        //
        val occurrenceMap = hashSetOf<Int>()

        for (num in nums) {
            if (occurrenceMap.contains(num)) return true
            occurrenceMap.add(num)
        }

        return false
    }
}