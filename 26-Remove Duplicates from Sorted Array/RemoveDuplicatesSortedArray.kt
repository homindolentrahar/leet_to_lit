fun main() {
    val nums = intArrayOf(1, 2, 2)
    val result = RemoveDuplicatesSortedArray.removeDuplicates(nums)

    println(result)
}

abstract class RemoveDuplicatesSortedArray {
    companion object {
        fun removeDuplicates(nums: IntArray): Int {
            // First attempt Mar 24, 2024 20:13
            // Runtime 269ms, Memory 43.1MB
            //
            // The first element always considered unique
//            var replaceIndex = 1
//
//            for (index in 1..<nums.size) {
//                // Check if current value NOT matches the previous value
//                if (nums[index - 1] != nums[index]) {
//                    // Replace item with index [replaceIndex] with current value
//                    nums[replaceIndex] = nums[index]
//                    // Increase the [replaceIndex] to execute next item
//                    replaceIndex++
//                }
//            }
//
//            println("Nums: ${nums.toList()}")
//
//            return replaceIndex
            //
            // Second attempt Mar 24, 2024 21:56
            // Runtime 213ms, Memory 39.9MB
            //
            var replaceIndex = 0
            var currentIndex = 1

            while (currentIndex < nums.size) {
                if (nums[currentIndex] == nums[replaceIndex]) {
                    currentIndex++
                    continue
                }

                replaceIndex++
                nums[replaceIndex] = nums[currentIndex]

                currentIndex++
            }

            return replaceIndex + 1
        }
    }
}
