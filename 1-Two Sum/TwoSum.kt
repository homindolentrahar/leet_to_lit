fun main() {
    val target = 9
    val nums = intArrayOf(2, 7, 11, 15)
    val result = Solution().twoSum(nums, target)

    println(result.toList())
}

class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        // First attempt Mar 15, 2024 05:56
        // Runtime 200ms, Memory 40.19MB
        //
//        val hashTemp = hashMapOf<Int, Int>()
//        var temp = intArrayOf()
//
//        for ((index, value) in nums.withIndex()) {
//            val diffs = target - value
//
//            if (hashTemp.contains(diffs)) {
//                temp += hashTemp[diffs] ?: 0
//                temp += index
//            } else {
//                hashTemp[value] = index
//            }
//        }
//
//        return temp
        //
        // Second attempt Apr 16, 2024 21:44
        // Runtime 188ms, Memory 38.08MB
        //
//        val hashTemp = hashMapOf<Int, Int>()
//        var temp = intArrayOf()
//
//        for (i in nums.indices) {
//            val diffs = target - nums[i]
//
//            if (hashTemp.keys.contains(diffs)) {
//                temp = intArrayOf(hashTemp[diffs] ?: 0, i)
//            }
//
//            hashTemp[nums[i]] = i
//        }
//
//        return temp
        //
        var indexes = intArrayOf()
        for (i in nums.indices) {
            for (j in i + 1..<nums.size) {
                if (nums[i] + nums[j] == target) {
                    indexes = intArrayOf(i, j)
                }
            }
        }

        return indexes
    }
}