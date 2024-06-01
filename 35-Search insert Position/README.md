# 35. Search Insert Position
#data-structure #algorithm #leetcode 
#### Level
#leetcode-easy 
#### Topics
#array #binary-search
#### Description
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
You must write an algorithm with `O(log n)` runtime complexity.
#### Definition
Example 1:
```
Input: nums = [1,3,5,6], target = 5
Output: 2
```
Example 2:
```
Input: nums = [1,3,5,6], target = 7
Output: 4
```
Constrains:
- `1 <= nums.length <= 104`
- `-104 <= nums[i] <= 104`
- `nums` contains **distinct** values sorted in **ascending** order.
- `-104 <= target <= 104`
#### Approaches
From the [[#Definition]] and [[#Description]], we know that:
1. `nums` contains **distinct** values sorted in **ASCENDING** order
2. We don't have to check every value within the `nums`
3. You have to write an algorithm with `O(log n)` runtime complexity
	Means you should use [[Binary Search]] to solve the problem.
	> 💡You could use approach with `O(n)` runtime complexity, but the `O(log n)` is more efficient
##### Binary Search
1. Define two pointers to track the start and end position of the `nums`
	- Initially assign first pointer (we will define as `start`) to 0
	- Initially assign second pointer (we will define as `end`) to the last index of `nums`
2. Create a while loop to iterate `nums`
	Create a while loop with condition if `start <= end`. Why? Because we want the loop to keep run if the `start` index is not reaching the `end` index, which means there's still chunks of `nums` that can be split into half
3. Define pointer that indicate the middle position of the `nums` (we will define as `mid`)
	You can define `mid` by looking for a median of the `nums` (both `(start + end) / 2` and `start + (end - start) / 2` works)
4. Compare the element at position `mid` with the `target`
	- If the element is **EQUAL** to `target`, return `mid` (as a position within `num`)
	- If the element are **GREATER** than `target`, shift the `end` pointer to position **BEFORE** `mid` (`mid - 1`). We want to remove the right chunk of `nums` (which is also **GREATER** than `target`)
	- If the element are **LESSER** than `target`, shift the `start` pointer to position **AFTER** `mid` (`mid + 1`). We want to remove the left chunk of `nums` (which is also **LESSER** than `target`)
5. Return the `start` to indicate the position or index where the `target` supposed to be inserted
#### Solution
- Kotlin
```kotlin
fun searchInsert(nums: IntArray, target: Int): Int {  
    var start = 0  
    var end = nums.size - 1  
  
    while (start <= end) {  
        val med = start + (end - start) / 2  
        
		if (target == nums[med]){
			return med
		}
  
        if (target > nums[med]) {  
            start = med + 1  
        } else {  
            end = med - 1  
        }  
    }  
    return start  
}
```
- Go
```go
func searchInsert(nums []int, target int) int {  
	start := 0
    end := len(nums) - 1  
  
    for start <= end {  
       mid := start + (end-start)/2  
  
       if nums[mid] == target {  
          return mid  
       }  
  
       if nums[mid] > target {  
          end = mid - 1  
       } else {  
          start = mid + 1  
       }  
    }  
    return start  
}
```
- Dart
```dart
int searchInsert(List<int> nums, int target) {
	int start = 0;  
	int end = nums.length - 1;  
  
	while (start <= end) {  
	  final med = start + (end - start) ~/ 2;
	  
	  if (target == nums[med]){
		  return med;
	  }
  
	  if (target > nums[med]) {  
	    start = med + 1;  
	  } else {  
	    end = med - 1;  
	  }}  
  
	return start;
}
```
#### References
- [AlgoEngine](https://www.youtube.com/watch?v=v4J_AWp-6EQ)
- [NeetCode](https://www.youtube.com/watch?v=K-RYzDZkzCI)