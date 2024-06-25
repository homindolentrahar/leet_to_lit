fun main() {
    val num = 5

    println(PascalTriangle.generate(num).toString())
}

object PascalTriangle {
    fun generate(numRows: Int): List<List<Int>> {
        //
        // First attempt 25 June 2024 12:18
        // Runtime: 163ms, Memory: 36.43MB
        // For loop
        //
//        val res = mutableListOf(mutableListOf(1))
//
//        for (i in 0..<numRows - 1) {
//            val lastRes = res[res.size - 1]
//            val temp = mutableListOf<Int>()
//            val row = mutableListOf<Int>()
//
//            temp.add(0, 0)
//            temp.addAll(1, lastRes)
//            temp.add(0)
//
//            for (j in 0..lastRes.size) {
//                row.add(temp[j] + temp[j + 1])
//            }
//
//            res.add(row)
//        }
//
//        return res
        //
        // Second attempt 25 June 2024 12:27
        // Runtime: 170ms, Memory: 37.16MB
        // Combinatoral Factorial
        //
//        val res = mutableListOf(mutableListOf(1))
//
//        for (i in 1..<numRows) {
//            val prevRow = res[i - 1]
//            val row = mutableListOf<Int>()
//
//            row.add(1)
//            for (j in 1..<i) {
//                row.add(prevRow[j - 1] + prevRow[j])
//            }
//            row.add(1)
//            res.add(row)
//        }
//
//        return res
        //
        // Third attempt 25 June 12:33
        // Runtime: 152ms, Memory: 36.69MB
        // Dynamic Programming with 1D Array
        //
        if (numRows == 0) {
            return listOf(emptyList())
        }
        if (numRows == 1) {
            return listOf(listOf(1))
        }

        val res = generate(numRows - 1).toMutableList()
        val prevRow = res[numRows - 2]
        val row = mutableListOf<Int>()

        row.add(1)
        for (i in 1..<numRows - 1) {
            row.add(prevRow[i - 1] + prevRow[i])
        }
        row.add(1)
        res.add(row)

        return res
    }
}