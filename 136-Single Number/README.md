# 136. Single Number

#algorithm #data-structure #leetcode

#### Level

#leetcode-easy

#### Topics

#array #bit-manipulation

#### Description

Given a**non-empty**array of integers`nums`, every element appears_twice_except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

#### Definition

Example 1:

```
Input: nums = [2,2,1]
Output: 1
```

Example 2:

```
Input: nums = [4,1,2,1,2]
Output: 4
```

Example 3:

```
Input: nums = [1]
Output: 1
```

Constraints:

- `1 <= numRows <= 30`
- `1 <= nums.length <= 3 * 104`
- `-3 * 104 <= nums[i] <= 3 * 104`
- Each element in the array appears twice except for one element which appears only once.

#### Approaches

From the [[#Definition]] and [[#Description]], we know that:

1. There is only **ONE** unique number inside `nums`
2. Each element appears twice, except the unique one
3. Implement solution with linear runtime complexity `O(n)`
4. Implement solution with no extra memory complexity `O(1)` (No Hash Map or Hash Set, because the memory complexity
   would be `O(n)`)
5. Use bit manipulation approach to achieve `O(1)` memory complexity (Use XOR bitwise operator)
   Basically, you would treat the number in a binary form (`1` and `0`), and use XOR bitwise operator to determine
   whether the number appears twice or once. XOR would return `1` if the two bits are different, otherwise return `0` if
   two bits are same.
   ![Binary illustration](https://github.com/homindolentrahar/leet_to_lit/blob/master/136-Single%20Number/assets/single_number.png)\
   *Credit to [NeetCode](https://www.youtube.com/@NeetCode) for image in
   this [video](https://youtu.be/qMPX1AOa83k?si=mxg4SFdl2dvU5v3a)*

   We'd sure that the bits of output **WOULD BE THE SAME AS** the bits of unique number (or single number). *"Why'd we
   so sure?"* Because the rest numbers within `nums` is duplicate, except the one that we're looking for.
   If you do XOR operation with duplicate numbers, it always return `0`. And if you do XOR operation between `n`
   and `0`, it always return the `n` (`n` represents any number between `1` and `0`).
   Now, consider the single number as an `n`, and the other duplicate numbers as a `0`. If we use XOR operation on those
   numbers, it will yield `n` or single number as a result, based on XOR logic explained before. Hence we could do XOR
   operation throughout `nums`, and get the unique single number as a result.

   Keypoints:
    - XOR operation on the duplicate numbers **ALWAYS RETURN `0`**
    - XOR operation between `n` and `0` **ALWAYS RETURN** `n` as a result
    - The bits of the output is gonna be the **SAME** as the bits of the single number, because the bits of other
      duplicate numbers **ALWAYS** gonna be `0`.

##### Bit Manipulation

1. **(OPTIONAL)** Create a guard checking first
   Check if `nums` is only contains one element. If so, just return the first element right away. Because if `nums` only
   contains one element, that means that the first element is always a single number.
2. Create variable to holds output value, define it as `val`
3. Create for-loop that iterate over the length of `nums`
   Inside the loop, assign `val` with the result of XOR operation of `val` and `nums[i]` (current num inside `nums`)
4. Return `val`

#### Solution

- Kotlin

```kotlin
fun singleNumber(nums: IntArray): Int {
    var value = 0

    for (num in nums) {
        value = value xor num
    }

    return value
}
```

- Go

```go
func singleNumber(nums []int) int{
val := 0
for num := range nums {
val = val ^ num
}

return val
}
```

- Dart

```dart
int singleNumber(List<int> nums) {
  int val = 0;

  for (int i = 0; i < nums.length; i++) {
    val = val ^ nums[i];
  }
  return val;
}
```

#### References

- [NeetCode](https://youtu.be/qMPX1AOa83k?si=mxg4SFdl2dvU5v3a)
- [Nick White](https://youtu.be/eXl0HBm2RrA?si=vieXK0fvVGNVD05_)
- [Algo Engine](https://youtu.be/xQPxiajgZLY?si=pBP1R6jU3C8fWoSt)