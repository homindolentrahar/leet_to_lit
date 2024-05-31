fun main() {
    val nums = intArrayOf(3, 2, 2, 3)
    val result = RemoveElement.removeElement(nums, 3)

    println(result)
}

abstract class RemoveElement {
    companion object {
        //
        // First attempt Mar 25, 2024 04:34
        // Runtime 137, Memory 34.50MB
        //
        fun removeElement(nums: IntArray, `val`: Int): Int {
            var index = 0
            for (i in nums.indices) {
                if (nums[i] != `val`) {
                    nums[index] = nums[i]
                    index++
                }
            }

            return index
        }
    }
}

