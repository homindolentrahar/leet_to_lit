package main

fun main() {
    val rowIndex = 4
    println(PascalTriangle2.getRow(rowIndex))
}

object PascalTriangle2 {
    fun getRow(rowIndex: Int): List<Int> {
        //
        // First attempt 26 June 2024 12:59
        // Runtime: 143ms, Memory: 34.76MB
        //
//        val res = mutableListOf(1)
//        var prev = 1
//
//        for (i in 1..rowIndex) {
//            val curr = (prev.toLong() * (rowIndex - i + 1) / i).toInt()
//            res.add(curr)
//            prev = curr
//        }
//
//        return res
        //
        // Second attempt 26 June 2024 13:03
        // Runtime: 140ms, Memory: 34.20MB
        // Recursive
        //
        if (rowIndex == 0) return listOf(1)

        val prevRow = getRow(rowIndex - 1)
        val row = mutableListOf<Int>()

        row.add(1)
        for (i in 1..<rowIndex) {
            row.add(prevRow[i] + prevRow[i - 1])
        }
        row.add(1)

        return row
    }
}

