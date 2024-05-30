#data-structure #algorithm #leetcode 
#### Level
#leetcode-easy 
#### Topics
#array #two-pointers
#### Description
Given an integer array `nums` sorted in **non-decreasing order**, remove the duplicates [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm) such that each unique element appears only **once**. The **relative order** of the elements should be kept the **same**. Then return _the number of unique elements in_ `nums`.
Consider the number of unique elements of `nums` to be `k`, to get accepted, you need to do the following things:
-  Change the array `nums` such that the first `k` elements of `nums` contain the unique elements in the order they were present in `nums` initially. The remaining elements of `nums` are not important as well as the size of `nums`.
- Return `k`.

> ❌ Cannot use built-in `pop()` method
#### Definition
Example 1:
```
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]

Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```
Example 2:
```
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```
Constrains:
-  `1 <= nums.length <= 3 * 104`
- `-100 <= nums[i] <= 100`
- `nums` is sorted in **non-decreasing** order.
#### Approaches
From the [[#Definition]], we know that:
1. List is sorted in **ASCENDING** order
2. You cannot use built-in `pop()` method to remove the duplicate element
	Using `pop()` method will shift all the elements backwards, so it's inefficient
3. You can start at the second element (index 1)
	Because the first element is **ALWAYS** unique within the array
4. You don't have to remove all the duplicates element
	You can modify the array, by replacing an element after the last unique element, with the next unique element
5. You don't have to be worry about the order of rest elements after the sorted one
##### Two Pointers
1. Define the first pointer that indicate the position of an element after the last unique element
	Since we already knew that the first element **ALWAYS** unique, then we start at index `1`
2. Create a loop that iterate the array, but start at index 1 (Second element)
	Since we already knew that the first element **ALWAYS** unique, we don't need to iterate from the first index to compare it with the previous element (Well you can, because it has no previous element)
3. Compare the current element with the previous element
4. If both elements are not equal (which means **UNIQUE**, then replace the element at index of first pointer, with the current element. Then increment the first pointer to the new position (after the last unique element)
5. Return the first pointer
#### Solution
- Kotlin
```kotlin
fun removeDuplicates(nums: IntArray): Int {  
    var replaceIndex = 0  
    var currentIndex = 1  
  
    while (currentIndex < nums.size) {  
	    if (nums[currentIndex] == nums[replaceIndex]) {  
            currentIndex++  
            continue  
        }  
  
        replaceIndex++  
        nums[replaceIndex] = nums[currentIndex]  
  
        currentIndex++  
    }  
  
    return replaceIndex + 1  
}
```
- Go
```go
func removeDuplicates(nums []int) int {  
    unique := 1  
  
    for i := 1; i < len(nums); i++ {  
       if nums[i] != nums[i-1] {  
          nums[unique] = nums[i]  
          unique++  
       }  
    }  
    return unique  
}
```
#### References
- [AlgoEngine](https://www.youtube.com/watch?v=oMr9lehS7Us)