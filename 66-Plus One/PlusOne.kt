fun main() {
    val digits = intArrayOf(9, 9)
    val result = PlusOne.plusOne(digits)

    println(result.joinToString())
}

abstract class PlusOne {
    companion object {
        fun plusOne(digits: IntArray): IntArray {
            //
            // First attempt Mar 27 2024 05:08
            // Runtime: 195ms, Memory: 38.10MB
            //
//            var tempDigits = digits
//
//            for (index in tempDigits.size - 1 downTo 0) {
//                if (tempDigits[index] == 9) {
//                    tempDigits[index] = 0
//                } else {
//                    tempDigits[index] += 1
//                    break
//                }
//            }
//
//            if (tempDigits[0] == 0) {
//                tempDigits = intArrayOf(1, *tempDigits)
//            }
//
//            println(tempDigits.toList())
//
//            return tempDigits
            //
            // Second attempt Mar 27 2024 05:26
            // Runtime: 135ms, Memory: 34.68MB
            //
            for (index in (digits.size - 1) downTo 0) {
                if (digits[index] < 9) {
                    digits[index]++
                    return digits
                }

                digits[index] = 0
            }

            val newDigits = IntArray(digits.size + 1)
            newDigits[0] = 1

            return newDigits
        }
    }
}
