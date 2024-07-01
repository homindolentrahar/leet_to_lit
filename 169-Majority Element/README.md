# 169. Majority Element

#algorithm #data-structure #leetcode

#### Level

#leetcode-easy

#### Topics

#array #hash-table #divide-conquer #sorting #counting

#### Description

Given an array`nums`of size`n`, return_the majority element_.
The majority element is the element that appears more than`⌊n / 2⌋`times. You may assume that the majority element
always exists in the array.

#### Definition

Example 1:

```
Input: nums = [3,2,3]
Output: 3
```

Example 2:

```
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

Constraints:

- `n == nums.length`
- `1 <= n <= 5 * 104`
- `-109 <= nums[i] <= 109`

#### Approaches

From the [[#Definition]] and [[#Description]], we know that:

1. The majority element is element that occurs more than `n/2` times
2. Assume that the majority element **ALWAYS** exists in the array
3. Follow-up question to solve the problem in linear time `O(n)` and constant space `O(1)`
4. You could use `HashMap` approach, but it will use `O(n)` space complexity. So consider using the next approach
5. We will be using [[Boyer-Moore Voting Algorithm]], since it tackle the follow-up question to solve this problem in
   linear time `O(n)` and constant space `O(1)`
   This algorithm works by utilize the vote count of the candidate. The candidate that has most vote count, become the
   candidate for the majority element before verification process. For more detail about this algorithm,
   refer [here](https://www.geeksforgeeks.org/boyer-moore-majority-voting-algorithm/)
   > Notes: This algorithm works, **ONLY** when the set of elements is guaranteed has majority element

##### HashMap Approach (NOT RECOMMENDED)

1. Create an empty `HashMap`, that will hold each element as `key` and the total occurrences as `value`, define it
   as `hashTable`
2. Create variable to hold the candidate of majority element, define it as `candidate`
3. Create variable to hold the max count of occurrences of each element, define it as `maxCount`
4. Create a for loop that iterate through `nums` array, define index as `i`
    1. Add the value of previous `hashTable[nums[i]]` by `1`
       We want to add each occurrences of `nums[i]` by `1`, hence we add `1` to the `hashTable[nums[i]]`
    2. Check if the value of `hashTable[nums[i]]` is **GREATER THAN** `maxCount`
       This means that `nums[i]` has more occurrences than the `maxCount` or the maximal count of occurrence's number.
       Hence, assign the `candidate` with `nums[i]`, since we consider `nums[i]` has the most occurrences within
       the `nums`
    3. Update the `maxCount` with the maximum value between `maxCount`'s current value and the value
       of `hashTable[element]`
       The `maxCount` will hold the maximum value between those two values
5. Return `candidate`

##### Boyer-Moore Voting Algorithm

Cases:

- Find the majority number candidate first
  > You can skipped this case, since the problem guarantee that `nums` always has majority number. But if you want to
  know how to find the majority element inside `nums`, you can follow this case

    1. Create variable that holds the candidate of majority number, define it as `candidate` and you can initialized
       with `-1`
    2. Create variable that holds the 'possibility' of the candidate becoming the majority number, define it as `vote`
    3. Create a loop that iterate through `nums`, define the index as `index`
        1. Check if `vote` equals to `0`
           If it's true, it means that we want to change the candidate number with the current number (`nums[index]`).
           We want to change the candidate because the `vote` is `0`, which means the previous candidate doesn't have
           the possibility to become majority number, thus we change the candidate with current number. Since we
           initialize the `candidate` with `-1` (which not a part of `nums`), we assign the first element of `nums`
           to `candidate`.
           Then increment `vote` by `1`, indicating that the new candidate has `1` occurrence
        2. Otherwise, check if the `candidate` is equal to current number (`nums[index]`)
            1. If it's true, it means that `candidate` has more occurrence, thus increment `vote` by `1`
            2. Otherwise, decrement `vote` by `1`
    4. Create variable that holds the total occurrences of found `candidate` within `nums`, define it as `count`
    5. Create another for loop to count the total occurrences of `candidate`
       If the current number (`nums[index]`) equals to found `candidate`, increment `count` by `1`
    6. Check if `count` is greater than `len(nums)/2`
       If it's true, it means that `candidate` occurs more than half of total `nums` times, thus return the `candidate`.
       Otherwise, you can return `-1`, which means that there is no majority number found (Although this is impossible,
       remembering the problem's constraint)
- Majority number candidate existent is guaranteed
    1. Create a variable to holds the candidate of majority number, define it as `candidate`
    2. Create a variable to holds the 'possibility' of the candidate becoming the majority number, define it as `count`
    3. Create a for loop that iterate through `nums`, define the index as `index`
        1. Check if the `count` equals to `0`
           If it's true, it means that the current candidate doesn't have possibility to become majority number
           within `nums`. So we change the candidate of majority number with the current number (`nums[index]`)
        2. Check if the `candiate` equals to current number (`nums[index]`)
           If it's true, it means that `candidate` occurs again within `nums`, thus increment the possibility (`count`)
           by `1`
           Otherwise, decrement the possibility (`count`) by `1`
    4. Return `candidate`

#### Solution

- Kotlin

```kotlin
fun majorityElement(nums: IntArray): Int {
    var candidate = 0
    var count = 0

    for (num in nums) {
        if (count == 0) {
            candidate = num
        }

        if (candidate == num) {
            count++
        } else {
            count--
        }
    }

    return candidate
}
```

- Go

```go
func majorityElement(nums []int) int {
candidate, count := 0, 0

for index := range nums {
if count == 0 {
candidate = nums[index]
}
if candidate == nums[index] {
count++
} else {
count--
}  
}

return candidate
}
```

- Dart

```dart
int majorityElement(List<int> nums) {
  int candidate = 0;
  int count = 0;

  for (int i = 0; i < nums.length; i++) {
    if (count == 0) {
      candidate = nums[i];
    }

    if (candidate == nums[i]) {
      count++;
    } else {
      count--;
    }
  }

  return candidate;
}
```

#### References

- [NeetCode](https://youtu.be/7pnhv842keE?si=wG4UIbbxkQBYLH3v)