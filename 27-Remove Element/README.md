# 27. Remove Element
#data-structure #algorithm #leetcode 
#### Level
#leetcode-easy 
#### Topics
#array #two-pointers 
#### Description
Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in `nums` [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm). The order of the elements may be changed. Then return _the number of elements in_ `nums` _which are not equal to_ `val`.
Consider the number of elements in `nums` which are not equal to `val` be `k`, to get accepted, you need to do the following things:
- Change the array `nums` such that the first `k` elements of `nums` contain the elements which are not equal to `val`. The remaining elements of `nums` are not important as well as the size of `nums`.
- Return `k`.
#### Definition
Example 1:
```
Input: nums = [3,2,2,3], val = 3
Output: 2, nums = [2,2,_,_]

Explanation: Your function should return k = 2, with the first two elements of nums being 2.
It does not matter what you leave beyond the returned k (hence they are underscores).
```
Example 2:
```
Input: nums = [0,1,2,2,3,0,4,2], val = 2
Output: 5, nums = [0,1,4,0,3,_,_,_]

Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
Note that the five elements can be returned in any order.
It does not matter what you leave beyond the returned k (hence they are underscores).
```
Constrains:
- `0 <= nums.length <= 100`
- `0 <= nums[i] <= 50`
- `0 <= val <= 100`
#### Approaches
From the [[#Definition]] and [[#Description]], we know that:
1. The order of the elements **DOESN'T MATTER**
2. You don't have to 'really' remove expected `val` from the array
	Because you ended up either creating new array or shifting elements in the array (it's inefficient)
3. We only care about the element that **NOT EQUAL** to `val`
	We don't really care about the element that equal to `val`. Because we only want to put **non-equal** to the beginning of array
##### Two Pointers
1. Define pointer at `0` that keep track about how many non-equal elements to the `val`
	It will also indicate the position that can be overridden with the non-equal elements to the `val`
2.  Create a loop that iterate the array 
	`index` in this loop act as an second pointer
3. Compare the current element with the target `val`
4. If both elements are **NOT EQUAL**, then replace the element with position of pointer with the current element. Then increment the pointer by one
	This will bring the not-equal element to the front
5. Return the pointer
#### Solution
- Kotlin
```kotlin
fun removeElement(nums: IntArray, `val`: Int): Int {  
    var index = 0  
    for (i in nums.indices) {  
        if (nums[i] != `val`) {  
            nums[index] = nums[i]  
            index++  
        }  
    }  
    
    return index  
}
```
- Go
```go
func removeElement(nums []int, val int) int {  
    pointer := 0  
    for index := range nums {  
       if nums[index] != val {  
          nums[pointer] = nums[index]  
          pointer++  
       }  
    }  
  
    return pointer  
}
```
- Dart
```dart
int removeElement(List<int> nums, int val) {
	int pointer = 0;  
	
	for (var index = 0; index < nums.length; index++) {  
		if (nums[index] != val) {  
			nums[pointer] = nums[index];  
			pointer++;  
		}  
	}  
	
	return pointer;
}
```
#### References
- [AlgoEngine](https://www.youtube.com/watch?v=pGKDzt0gk-A)
- [NeetCode](https://www.youtube.com/watch?v=Pcd1ii9P9ZI)