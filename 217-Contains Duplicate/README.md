# 217. Contains Duplicate

#algorithm #data-structure #leetcode

#### Level

#leetcode-easy

#### Topics

#array #hash-table #sorting

#### Description

Given an integer array`nums`, return`true`if any value appears**at least twice**in the array, and return`false`if every
element is distinct.

#### Definition

Example 1:

```
Input: nums = [1,2,3,1]
Output: true
```

Example 2:

```
Input: nums = [1,2,3,4]
Output: false
```

Example 3:

```
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
```

Constraints:

- `1 <= nums.length <= 105`
- `-109 <= nums[i] <= 109`

#### Approaches

From the [[#Definition]] and [[#Description]], we know that:

1. There is value that will appears **AT LEAST TWICE** in the `nums`
2. Return `true` right away whenever the duplicate value appeared
3. By default, return `false` if there's no other duplicate value inside `nums`
4. You can use several approaches for solving this problem:
    - Brute-force (**NOT RECOMMENDED**)
      It's the most inefficient approach to solve this problem, because we have to compare each single number with every
      other number inside `nums`. This approach will take `O(n^2)` Time complexity (which is pretty slow), but they
      have `O(1)` Space complexity after all.
    - Sorting (**COULD BE IMPROVED**)
      We knew that we would return `true` and stop the program whenever we stumble upon duplicate value. Since it's
      number, we can sort `nums` to make duplicate values appears next to each other. By that, we can easily traverse
      through `nums` and check if the duplicates between current number and the next number. This approach will
      take `O(log^n)` Time complexity and `O(1)` Space complexity, because we don't create new array.
    - HashSet (**RECOMMENDED**)
      Basically, you create a `HashSet` that holds the value of each number inside `nums`. While traversing, you add
      each number to the `HashSet`. We would determine duplicate values by checking if the number is already exists
      inside the `HashSet`, and then return `true` right away. Although the Time complexity would be `O(n)`, but we have
      to sacrifice the Space complexity little bit, by `O(n)`, since we create a new `HashSett`.

##### HashSet

1. Create `HashSet` to holds each number inside `nums`, define it as `occurrenceMap`
2. Create for loop that iterate through `nums`, define index as `i`
    1. Check if current number (`nums[i]`) already exists inside `occurrenceMap`. If it is, then return `true`, since we
       know that number must be duplicate of previous number.
    2. Add current number (`nums[i]`) into the `occurrenceMap`
3. Return `false`

#### Solution

- Kotlin

```kotlin
fun containsDuplicate(nums: IntArray): Boolean {
    val occurrenceMap = hashSetOf<Int>()

    for (num in nums) {
        if (occurrenceMap.contains(num)) return true
        occurrenceMap.add(num)
    }

    return false
}
```

- Go

```go
func containsDuplicate(nums []int) bool {
occurrenceMap := make(map[int]bool)

for index := range nums {
if _, ok := occurrenceMap[nums[index]]; ok {
return true
}

occurrenceMap[nums[index]] = true  
}

return false
}
```

- Dart

```dart
bool containsDuplicate(List<int> nums) {
  final occurrenceMap = <int>{};

  for (int i = 0; i < nums.length; i++) {
    if (occurrenceMap.contains(nums[i])) {
      return true;
    }
    occurrenceMap.add(nums[i]);
  }

  return false;
}
```

#### References

- [NeetCode on YouTube](https://youtu.be/3OamzN90kPg?si=abJu05n1tuM1pzPU)
- [Nick White on YouTube](https://youtu.be/4oZsPXG9B94?si=4Mi7ydr2-2RNODMa)
- [Nikhil Lohia on YouTube](https://youtu.be/0LIctkgJ2hQ?si=FvI0YX8Eaf_sqSms)