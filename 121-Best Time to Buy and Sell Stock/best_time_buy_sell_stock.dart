import 'dart:math';

void main() {
  final prices = [7, 1, 5, 3, 6, 4];
  print(BestTimeBuySellStock.maxProfit(prices).toString());
}

abstract class BestTimeBuySellStock {
  static int maxProfit(List<int> prices) {
    //
    // First attempt 27 June 2024
    // Runtime: 410ms, Memory: 189.11MB
    // Time complexity: O(n^2), Space complexity: O(1)
    //
    int profit = 0;
    int buyPrice = prices[0];

    for (int i = 1; i < prices.length; i++) {
      if (prices[i] > buyPrice) {
        profit = max(profit, prices[i] - buyPrice);
      } else {
        buyPrice = prices[i];
      }
    }

    return profit;
  }
}
