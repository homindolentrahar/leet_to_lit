# 118. Pascal's Triangle

#algorithm #data-structure #leetcode

#### Level

#leetcode-easy

#### Topics

#array #dynamic-programming

#### Description

![Pascal Triangle](https://github.com/homindolentrahar/leet_to_lit/blob/master/119-Pascal's%20Triangle%20II/assets/pascal_triangle.gif)\
Given an integer`rowIndex`, return the`rowIndexth`(**0-indexed**) row of the**Pascal's triangle**.
In**Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:

#### Definition

Example 1:

```
Input: rowIndex = 3
Output: [1,3,3,1]
```

Example 2:

```
Input: rowIndex = 0
Output: [1]
```

Example 3:

```
Input: rowIndex = 1
Output: [1,1]
```

Constraints:

- `0 <= numRows <= 33`

#### Approaches

From the [[#Definition]] and [[#Description]], we know that:

1. `rowIndex` start at `0`
2. We need to return the row of Pascal Triangle with index `rowIndex`
3. The number of values inside a row is **ALWAYS EQUALS** to the number of that row
   For example, row 1 will only have 1 values, row 4 will have 4 values, etc.
4. Each number inside a row, is a sum of the two numbers directly above its row
   See the GIF above for more clear example
5. We can use several approaches to solve this problem as shown below

##### Nested Loop - O(n^2) Time complexity

1. Create a single array that has initial element of `[1]`
   We already knew that the first row of Pascal Triangle is always `1` for every row, so we initialized the array
   with `[1]`. We can define the array as `res`.
2. Create a for loop that iterate `rowIndex` times, with index defined as `i`. (`0 <= i < rowIndex`)
   We want to loop for `rowIndex` times. We already knew that the length of a row is always the same as the position of
   that row. (Row 2 has 2 elements, Row 4 has 4 elements, etc.). So we want to iterate as much as `rowIndex`.
    1. Create an empty array with the length of `len(res) + 1`, defined as `row`
       We want to make empty array with that length, because we know that the current row has more length than the
       previous one.
    2. Create another for loop that iterate `len(res)` times, with index defined as `j`
       We want to take the element from previous row, then calculated and added to the current `row`.
        - Assign `row[j]` with `row[j] + res[j]`
        - Assign `row[j]` with `row[j + 1] + res[j]`
          "*Why we'd do that?*"
          See, the concept of Pascal Triangle is that each value inside a row has some kind of "2 children" int the next
          row (Refer to this image below). We treat a value as a "parent" that has "2 children" below it. So, we can
          populate the value of the "parent" into the "2 children" below it. But, because there is a "child" that
          belongs to "2 parents", we sum the value of "child" (that has been populated by the first "parent" before)
          with the value of second "parent" (`4` is the sum of `1` (first "parent") and `3` (second "parent")) . Hence,
          we have to do two points above.
          ![[Pasted image 20240626133732.png]]\
          ![Solution Illustration](https://github.com/homindolentrahar/leet_to_lit/blob/master/119-Pascal's%20Triangle%20II/assets/solution_illustration.png)\
          *Credit to [NeetCode](https://www.youtube.com/@NeetCode) for image in
          this [video](https://youtu.be/k1DNTyal77I?si=_7MrqrBDjMlRTMP8)*
    3. Assign `res` with `row`
3. Return `res`

##### Recursive

1. Create a guard checking first (`rowIndex == 0`)
   First, we need to check if `rowIndex` is equal to `0`. If its equal, return `[1]` right away
2. Create an array that holds value of recursive `getRow(rowIndex - 1)` function's call, define it as `previousRow`
   We call the `getRow()` function itself to create recursive call. We pass `rowIndex - 1` as a parameter, because we
   want to holds the value of the previous row.
3. Create an empty array that will hold the current row, define it as `row`
4. Append `row` with `1` at the start of the current row
   We knew that in every row of Pascal Triangle starts and ends with `1`, so we append it first.
5. Create for loop that iterate the `rowIndex` and start at index `1`, with index defined as `i`
   We start at index `1` because we already define the `0` index or first value as `1` before.
   Inside the loop, append the `row` with the result of sum `prevRow[i-1]` and `prevRow[j]`. We want to sum two numbers
   directly above current row. But because we start at index `1`, we use `j-1` as the first value, because it points on
   the first value or index `0`.
6. Append `row` with `1` at the end of the current row
   We knew that in every row of Pascal Triangle starts and ends with `1`, so we append it at the last of `row`.
7. Return `row`

#### Solution

- Kotlin

```kotlin
// Recursive
fun getRow(rowIndex: Int): List<Int> {
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
```

- Go

```go
// Recursive
func getRow(rowIndex int) []int {
if rowIndex == 0 {
return []int{1}  
}

previousRow := getRow(rowIndex - 1)
var row []int

row = append(row, 1)
for i := 1; i < rowIndex; i++ {
row = append(row, previousRow[i-1]+previousRow[i])
}
row = append(row, 1)

return row
}

// Nested for loop
func getRow(rowIndex int) []int {
res := []int{1}

for range rowIndex {
row := make([]int, len(res)+1)

for j := range len(res) {
row[j] += res[j]
row[j+1] += res[j]
}

res = row
}

return res
}
```

- Dart

```dart
// Recursive
List<int> getRow(int rowIndex) {
  if (rowIndex == 0) return [1];

  final prevRow = getRow(rowIndex - 1);
  final row = <int>[];

  row.add(1);
  for (int i = 1; i < rowIndex; i++) {
    row.add(prevRow[i - 1] + prevRow[i]);
  }
  row.add(1);

  return row;
}
```

#### References

- [NeetCode](https://youtu.be/k1DNTyal77I?si=_7MrqrBDjMlRTMP8)
- [vanAmsem](https://youtu.be/BY95fxcrjag?si=ObE_zwS7huCHyol6)