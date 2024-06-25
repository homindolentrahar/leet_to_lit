# 118. Pascal's Triangle 

#algorithm #data-structure #leetcode

#### Level

#leetcode-easy

#### Topics

#array #dynamic-programming

#### Description

![Pascal Triangle](https://github.com/homindolentrahar/leet_to_lit/blob/master/118-Pascal's%20Triangle/assets/pascal_triangle.gif)
Given an integer`numRows`, return the first `numRows` of**Pascal's triangle**.
In**Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:

#### Definition

Example 1:

```
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```

Example 2:

```
Input: numRows = 1
Output: [[1]]
```

Constraints:

- `1 <= numRows <= 30`

#### Approaches

From the [[#Definition]] and [[#Description]], we know that:

1. `numRows` is always positive number
2. We need to return the triangle as a 2D Array
   The first dimension array represents the row of a triangle, and the second dimension array represents the values
   within the row.
3. The number of values inside a row is **ALWAYS EQUALS** to the number of that row
   For example, row 1 will only have 1 values, row 4 will have 4 values, etc.
4. Each number inside a row, is a sum of the two numbers directly above its row
   See the GIF above for more clear example
5. We can use several approaches to solve this problem as shown below

##### For Loop

1. Create a 2D array that has initial element of `[1]`
   We already knew that the first row of Pascal Triangle is always `1` for every `numRows`, so we initialized the array
   with `[1]`. We can define the array as `res`.
2. Create a loop that iterate `numRows - 1`
   We want to loop for `numRows - 1`  times. "*Why we should loop for `numRows - 1` times?*" Because we already get the
   first row of Pascal Triangle by the initialization above. So we don't need to loop for the first row, hence the
   number of loop decreased by `1`.
    1. Create variable to hold a value of the previous row in the triangle
       We can do this by accessing `res[len(res) - 1]`. We can define it as `prevRow`
    2. Create another variable to hold a value of the previous row in the triangle, but add `0` as prefix and suffix
       inside the row. We can define it as `temp`
       "*Why we should add `0` as a prefix and suffix inside a row?*" Referring to the statement from
       the [[#Description]], `Each number inside a row, is a sum of the two numbers directly above its row`. If you
       looked up to the GIF above, the first and last of every row is always `1`. By that statement, other values inside
       row is a result of sum. But the first and last value of row only has `1` as a number above it, there is no other
       number to calculate the sum.
       So, we kinda simulate that we have another number that will resulted in `1`, if we sum that number with `1`. What
       number is that? The answer is `0`. But don't worry, the so-called prefix and suffix number will not included in
       our Pascal Triangle later, just for the calculation.
    3. Create an empty array to hold values of the current row. We can define it as `row`
    4. Create another loop that iterate as much as the length of `prevRow` + `1`, with index defined as `j`
       We create another loop to assign values inside current row. "*Why we should loop for `len(prevRow) + 1` times?*"
       Remember that `Each number inside a row, is a sum of the two numbers directly above its row`, so we should
       iterate the `prevRow` in order to sum two numbers inside it. We add the length of `prevRow` by `1`, because we
       already add `0` as a suffix and prefix on `temp`. Otherwise, it will yield `IndexOutOfBound` or something.
       Inside the loop, append the `row` with the result of sum `temp[j]` and `temp[j+1]`. We want to sum two numbers
       that placed directly above current row, and add it to the current row.
3. Append the `res` with the `row`, we want to insert `row` that holds value of the current row into the triangle.
4. Return `res`

##### Combinatoral Formula

1. Create a 2D array that has initial element of `[1]`
   We already knew that the first row of Pascal Triangle is always `1` for every `numRows`, so we initialized the array
   with `[1]`. We can define the array as `res`.
2. Create a loop that iterate `numRows` and start at index `1`, with index defined as `i`
   We want to start at index `1` because we already defined the first row or `0` index at the initialization of `res`.
    1. Create variable to hold a value of the previous row in the triangle
       We can do this by accessing `res[i - 1]`. We can define it as `prevRow`
    2. Create an empty array to hold values of the current row. We can define it as `row`
    3. Append `row` with `1` at the start of the current row
       We knew that in every row of Pascal Triangle starts and ends with `1`, so we append it first.
    4. Create another loop that iterate the `i` and start at index `1`, with index defined as `j`
       We start at index `1` because we already define the `0` index or first value as `1` before.
       Inside the loop, append the `row` with the result of sum `prevRow[j-1]` and `prevRow[j]`. Same as before, we want
       to sum two numbers directly above current row. But because we start at index `1`, we use `j-1` as the first
       value, because it points on the first value or index `0`.
    5. Append `row` with `1` at the end of the current row
       We knew that in every row of Pascal Triangle starts and ends with `1`, so we append it at the last of `row`.
3. Append the `res` with the `row`, we want to insert `row` that holds value of the current row into the triangle.
4. Return `res`

##### Dynamic Programming with 1D Array

1. Create initial checking for recursion purpose
    - If `numRows` equals to `0`, return the empty 2D Array
    - if `numRows` equals to `1`, return 2D Array with initial value of `[1]`
2. Create a 2D array that holds value of recursive `generate(numRows - 1)` function's call
   We call the `generate()` function itself to create recursive call. We pass `numRows - 1` as a parameter, because we
   want to holds the value of the last triangle before `numRows`.
3. Create variable to hold the previous row inside the last triangle
   We can define it as `prevRow`, and assigned it with `res[numRows - 2]`. "*Why we accessing `numRows - 2`?*" Because
   we already define the value of first and second row in the 1st step. Because we call a recursive function, it will
   run the 1st step again an we will get those values. So we start at the third row or index `2`.
4. Create empty array to hold values of the current row, define it as `row`
5. Append `row` with `1` at the start of the current row
   We knew that in every row of Pascal Triangle starts and ends with `1`, so we append it first.
6. Create a loop that iterate `numRows - 1` and start at index `1`, with index defined as `i`
   We want to start at index `1` because we already define the `0` index or first value as `1` before. We subtract the
   length of `prevRow` by `1`, because we start at index `2` or third element. Otherwise, it will
   yield `IndexOutOfBound` or something.
   Inside the loop, append the `row` with the result of sum `prevRow[j-1]` and `prevRow[j]`. Same as before, we want to
   sum two numbers directly above current row. But because we start at index `1`, we use `j-1` as the first value,
   because it points on the first value or index `0`.
7. Append `row` with `1` at the end of the current row
   We knew that in every row of Pascal Triangle starts and ends with `1`, so we append it at the last of `row`.
8. Append the `res` with the `row`, we want to insert `row` that holds value of the current row into the triangle.
9. Return `res`

#### Solution

- Kotlin

```kotlin
// For loop
fun generate(numRows: Int): List<List<Int>> {
    val res = mutableListOf(mutableListOf(1))

    for (i in 0..<numRows) {
        val lastRes = res[res.size - 1]
        val temp = mutableListOf<Int>()
        val row = mutableListOf<Int>()

        temp.add(0, 0)
        temp.addAll(1, lastRes)
        temp.add(0)

        for (j in 0..lastRes.size) {
            row.add(temp[j] + temp[j + 1])
        }
        res.add(row)
    }

    return res
}

// Combinatoral Factorial
fun generate(numRows: Int): List<List<Int>> {
    val res = mutableListOf(mutableListOf(1))

    for (i in 1..<numRows) {
        val prevRow = res[i - 1]
        val row = mutableListOf<Int>()

        row.add(1)
        for (j in 1..<i) {
            row.add(prevRow[j - 1] + prevRow[j])
        } row . add (1)
        res.add(row)
    }

    return res
}

// Dynamic Programming with 1D Array
fun generate(numRows: Int): List<List<Int>> {
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
```

- Go

```go
// For loop
func generate(numRows int) [][]int{
res := [][]int{{1}}

for range numRows - 1 {
lastRes := res[len(res)-1]
var temp = []int{0}
var row []int

temp = append(temp, lastRes...)
temp = append(temp, 0)

for j := range len(lastRes) + 1 {
row = append(row, temp[j]+temp[j+1])
}

res = append(res, row)
}

return res
}

// Combinatoral Factorial
func generate(numRows int) [][]int{
res := [][]int{{1}}

for i := 1; i < numRows; i++ {
prevRow := res[i-1]
var row []int
row = append(row, 1)

for j := 1; j < i; j++ {
row = append(row, prevRow[j-1]+prevRow[j])
}

row = append(row, 1)
res = append(res, row)
}

return res
}

// Dynamic Progamming with 1D Array
func generate(numRows int) [][]int{
if numRows == 0 {
return [][]int{}  
}
if numRows == 1 {
return [][]int{{1}}  
}

res := generate(numRows - 1)
prevRow := res[numRows-2]
var row []int

row = append(row, 1)
for i := 1; i < numRows-1; i++ {
row = append(row, prevRow[i-1]+prevRow[i])
}
row = append(row, 1)
res = append(res, row)

return res
}
```

- Dart

```dart
// For loop
List<List<int>> generate(int numRows) {
  final res = [[1]];

  for (int i = 0; i < numRows - 1; i++) {
    final lastRes = res[res.length - 1];
    final temp = [0, ...lastRes, 0];
    final row = <int>[];

    for (int j = 0; j < lastRes.length + 1; j++) {
      row.add(temp[j] + temp[j + 1]);
    }
    res.add(row);
  }

  return res;
}

// Combinatoral Factorial
List<List<int>> generate(int numRows) {
  final res = [[1]];

  for (int i = 0; i < numRows; i++) {
    final prevRows = res[i - 1];
    final row = <int>[];

    row.add(1);
    for (int j = 1; j < i; j++) {
      row.add(prevRows[j - 1] + prevRows[j]);
    }
    row.add(1);
    res.add(row);
  }

  return res;
}

// Dynamic Programming with 1D Array
List<List<int>> generate(int numRows) {
  if (numRows == 0) {
    return [[]];
  }

  if (numRows == 1) {
    return [[1]];
  }

  final res = List.of(generate(numRows - 1));
  final prevRow = res[numRows - 2];
  final row = <int>[];

  row.add(1);
  for (int i = 1; i < numRows - 1; i++) {
    row.add(prevRow[i - 1] + prevRow[i]);
  }
  row.add(1);
  res.add(row);

  return res;
}
```

#### References

- [NeetCode](https://youtu.be/nPVEaB3AjUM?si=9U1z0mBZP-jcJHvD)
- [Nick White](https://youtu.be/icoql2WKmbA?si=3GNW191x50tvaHwf)