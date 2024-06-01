fun main() {
    val nums = intArrayOf(1, 3, 5, 6)
    val result = SearchInsertPosition.searchInsert(nums, 5)

    println(result)
}

abstract class SearchInsertPosition {
    companion object {
        fun searchInsert(nums: IntArray, target: Int): Int {
            //
            // First attempt Mar 29, 2024 04:48
            // Runtime 144ms, Memory 37.63MB
            //
            var start = 0
            var end = nums.size - 1

            while (start <= end) {
                val med = start + (end - start) / 2

                if (target == nums[med]) {
                    return med
                }

                if (target > nums[med]) {
                    start = med + 1
                } else {
                    end = med - 1
                }
            }

            return start
        }
    }
}
