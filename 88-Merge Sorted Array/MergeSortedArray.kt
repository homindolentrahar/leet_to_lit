fun main() {
    val nums1 = intArrayOf(1, 2, 3, 0, 0, 0)
    val nums2 = intArrayOf(2, 5, 6)
    val m = 3
    val n = 3

    MergeSortedArray.merge(nums1, m, nums2, n)
}

object MergeSortedArray {
    fun merge(nums1: IntArray, m: Int, nums2: IntArray, n: Int): Unit {
        //
        // First attempt 06 June 2024 15:26
        // Runtime: 198ms, Memory: 37.80MB
        //
//        var lastPointer = m + n - 1
//        var pointerM = m
//        var pointerN = n
//
//        while (pointerM > 0 && pointerN > 0) {
//            if (nums1[pointerM - 1] > nums2[pointerN - 1]) {
//                nums1[lastPointer] = nums1[pointerM - 1]
//                pointerM--
//            } else {
//                nums1[lastPointer] = nums2[pointerN - 1]
//                pointerN--
//            }
//
//            lastPointer--
//
//        }
//
//        while (pointerN > 0) {
//            nums1[lastPointer] = nums2[pointerN - 1]
//            pointerN--
//            lastPointer--
//        }
//
//        println(nums1.joinToString(", "))
        //
        // Third attempt 06 June 2024 15:32
        // Runtime: 159ms, Memory: 35.45MB
        //
        var lastPointer = m + n - 1
        var pointerM = m - 1
        var pointerN = n - 1

        while (pointerM >= 0 && pointerN >= 0) {
            if (nums1[pointerM] > nums2[pointerN]) {
                nums1[lastPointer] = nums1[pointerM]
                pointerM--
            } else {
                nums1[lastPointer] = nums2[pointerN]
                pointerN--
            }

            lastPointer--

        }

        while (pointerN >= 0) {
            nums1[lastPointer] = nums2[pointerN]
            pointerN--
            lastPointer--
        }

        println(nums1.joinToString(", "))
    }
}

