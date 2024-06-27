package main

fun main() {
    val prices = intArrayOf(7, 1, 5, 3, 6, 4)
    println(BestTimeBuySellStock.maxProfit(prices))
}

object BestTimeBuySellStock {
    fun maxProfit(prices: IntArray): Int {
        //
        // First attempt 27 June 2024 15:45
        // Runtime: 467ms, Memory: 55.70MB
        // Time complexity: O(n), Space complexity: O(1)
        //
        var profit = 0
        var buyPrice = prices[0]

        for (i in 1..<prices.size) {
            if (prices[i] > buyPrice) {
                profit = maxOf(profit, prices[i] - buyPrice)
            } else {
                buyPrice = prices[i]
            }
        }

        return profit
    }
}

