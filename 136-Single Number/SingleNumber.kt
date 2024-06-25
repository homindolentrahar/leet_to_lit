fun main() {
    val nums = intArrayOf(2, 2, 1)

    println("Output: ${SingleNumber.singleNumber(nums)}")
}

object SingleNumber {
    fun singleNumber(nums: IntArray): Int {
        //
        // First attempt 25 June 2024 16:02
        // Runtime: 216ms, Memory: 39.59MB
        //
        var value = 0

        for (num in nums) {
            value = value xor num
        }

        return value
    }
}