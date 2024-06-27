# 121. Best Time to Buy and Sell Stock

#algorithm #data-structure #leetcode

#### Level

#leetcode-easy

#### Topics

#array #two-pointers

#### Description

You are given an array`prices`where`prices[i]`is the price of a given stock on the`ith`day.
You want to maximize your profit by choosing a**single day**to buy one stock and choosing a**different day in the future
**to sell that stock.
Return_the maximum profit you can achieve from this transaction_. If you cannot achieve any profit, return`0`.

#### Definition

Example 1:

```
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.

Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
```

Example 2:

```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```

Constraints:

- `1 <= prices.length <= 105`
- `0 <= prices[i] <= 104`

#### Approaches

From the [[#Definition]] and [[#Description]], we know that:

1. You have to choose "buy day" and "sell day" that will yield maximum profit
   In other words, the you have to choose "sell day" that has maximum profit (largest profit) when subtracted by "buy
   day". *"Selling low, buying high"*, they said.
2. You **CANNOT** choose "sell day" that occurred **BEFORE** "buy day"
   It won't make sense if you sell on `Monday`, but buy in `Wednesday`. The "buy day" **ALWAYS**** have to precede the "
   sell day". Unfortunately, you can't time travel to the past.
3. If the "buy day" has larger price than any other followed "sell day" price, return `0`
   It means that you won't gain a profit and no transactions are done
4. Don't take brute-force approach
   Brute-force approach will compare each day's price with other day's price one-by-one, which means your program will
   run in a nested loop. This approach is not effective, as it will have `O(n^2)` time complexity. Use other approach
   that has `O(n)` time complexity.

##### Two Pointers

1. Create variable to hold the value of maximum profit, define it as `profit`
2. Create variable that acts as a "pointer", pointing at the "buy price", define it as `buyPrice`
   This pointer will track the position of "buy day", the day we want to buy a stock. Assign this variable with the
   price of that "buy day". But first, initialized it as the first element of `prices` array, since we consider that we
   would buy at the first day initially.
3. Create a for loop that iterate through `prices` array, start with index `1` that defined as `i`
   With this loop, we will compare the price of "buy day" to each following price of "sell day" within `prices`. We
   start at index `1`, because the first element or index `0` would be a "buy day" initially, thus that day cannot be
   a "sell day".
   Inside the loop, we compare the price of current day (`prices[i]`) with the price of "buy day" (`buyPrice`)
    - If the price of current day (`prices[i]`) is **GREATER THAN** the price of "buy day", then assign `profit` with
      the profit
      We can calculate the profit by subtract the current price with the price we have tracked
      before (`prices[i] - buyPrice`).
      Then we need to compare it with the current value of `profit`. Why? Because we want to assign the **MAXIMUM PROFIT
      **, not just the profit. If we don't compare those two price and assign the profit directly, it will assign any
      profit to the variable `profit` as long as the current day's price is larger than the tracked price. We can
      use `max()` value to determine the **MAXIMUM PROFIT** of those two profits.
    - Or else, move the `buyPrice` pointer to the current day's price
      Since we encounter the price that lower than the price we tracked before, we will assign the lower price to the
      pointer (`buyPrice`)
4. Return `profit`

#### Solution

- Kotlin

```kotlin
fun maxProfit(prices: IntArray): Int {
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
```

- Go

```go
func maxProfit(prices []int) int {
profit := 0
buyPrice := prices[0]

for i := 1; i < len(prices); i++ {
if prices[i] > buyPrice {
profit = max(profit, prices[i] - buyPrice)
} else {
buyPrice = prices[i]
}    
}

return profit
}
```

- Dart

```dart
int maxProfit(List<int> prices) {
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
```

#### References

- [NeetCode](https://youtu.be/1pkOgXD63yU?si=rfY6btzbJdX6A2k-)
- [Algo Engine](https://youtu.be/ioFPBdChabY?si=pTKMmeH6R_b53qzd)