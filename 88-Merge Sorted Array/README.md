# 88. Merge Sorted Array
#data-structure #algorithm #leetcode 
#### Level
#leetcode-easy 
#### Topics
#array #two-pointers #sorting
#### Description
You are given two integer arrays `nums1` and `nums2`, sorted in **non-decreasing order**, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively.

**Merge** `nums1` and `nums2` into a single array sorted in **non-decreasing order**.

The final sorted array should not be returned by the function, but instead be _stored inside the array_ `nums1`. To accommodate this, `nums1` has a length of `m + n`, where the first `m` elements denote the elements that should be merged, and the last `n` elements are set to `0` and should be ignored. `nums2` has a length of `n`.
#### Definition
Example 1:
```
Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
Output: [1,2,2,3,5,6]

Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
```
Example 2:
```
Input: nums1 = [1], m = 1, nums2 = [], n = 0
Output: [1]

Explanation: The arrays we are merging are [1] and [].
The result of the merge is [1].
```
Example 3:
```
Input: nums1 = [0], m = 0, nums2 = [1], n = 1
Output: [1]

Explanation: The arrays we are merging are [] and [1].
The result of the merge is [1].
Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.
```
Constraints:
-  `nums1.length == m + n`
- `nums2.length == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-109 <= nums1[i], nums2[j] <= 109`
#### Approaches
From the [[#Definition]] and [[#Description]], we know that:
1. You cannot create a new array
	You can only alter and modify `nums1`. In other words, you need to merge `nums2` into `nums1`
2. We would start iterating from the last position of `nums1`
	Because we know that 'empty space' that can be replaced are sorted in the last of `nums1`
3. `m` represents the length of `nums1`, and `n` represents the length of `nums2`
4. We don't count `0` as an element or member of both `nums1` and `nums2`
	So if `nums1` or `nums2` has `0` as the **ONLY** element, then it will gives `0` to `m` and `n` respectively
5. Length of the final `nums1` is `m + n`
##### Two Pointers
![Illustration](https://i.loli.net/2019/05/14/5cdac3209479073662.gif)
1. Define the `lastPointer` to track the position of 'empty space' in `nums1`
	Assign it with `m + n - 1`, because we need to iterate `nums1` from the last position (since we knew that `nums1` sorted in a way so the 'empty space' appears on the last position)
2. Define the `while` loop
	Define a loop that runs while `m` and `n` is **GREATER** than `0`. We need to make sure that loop will run if `m` and `n` (as a pointer for `nums1` and `nums2`) does not reach the first position.
3. Compare the element of `nums1[m - 1]` with `nums2[n - 1]`
	We need to compare the elements at last position of each array, in order to determine the order of element `nums2` that will be merged into `nums1`.
	We decrease `m` and `n` by `1` because `m` and `n` is the length of `nums1` and `nums2`, and we need the last element's position of both arrays
	- If the element `nums1[m - 1]` is **GREATER** than element `nums2[n - 1]`, then replace the element of `nums1[lastPointer]` with `nums1[m - 1]`, and then decrease `m` by `1`
		This means that element of `nums1` is greater than element that will be merged from `nums2`. So we need to shift the element of `nums1` to the right, in order to merge element from `nums2` to the position of `lastPointer`. After that, we shift the `m` pointer to the left.
	- If the element `nums1[m - 1]` is **LESSER** than element `nums2[n - 1]`, then replace the element of `nums1[lastPointer]` with `nums2[n - 1]`, and then decrease `n` by `1`
		This means that element of `nums1` is lesser than element that will be merged from `nums2`. So we just need to replace the element at `lastPointer` with element from `nums2`. After that, we shift the `n` pointer to the left.
4. Shift the `lastPointer` to the left
5. Define another `while` loop to handle edge cases
	*"What if the first element of `nums1` is still lesser than element merged from `nums2`?"*. In order to handle that, wee need to create another loop that runs while `n` is **GREATER** than `0`. This means that `nums2` has larger element than `nums1`. 
	We just need to replace element at `lastPointer` in `nums1`, with the `nums2[n - 1]`, just like in the second case above. And then we need to shift pointer `n` and `lastPointer` to the left.
#### Solution
- Kotlin
```kotlin
fun merge(nums1: IntArray, m: Int, nums2: IntArray, n: Int): Unit {
	var lastPointer = m + n - 1  
	var pointerM = m - 1  
	var pointerN = n - 1  
  
	while (pointerM >= 0 && pointerN >= 0) {  
	    if (nums1[pointerM] > nums2[pointerN]) {  
	        nums1[lastPointer] = nums1[pointerM]  
	        pointerM--  
	    } else {  
	        nums1[lastPointer] = nums2[pointerN]  
	        pointerN--  
	    }  
  
	    lastPointer--  
	}  
  
	while (pointerN >= 0) {  
	    nums1[lastPointer] = nums2[pointerN]  
	    pointerN--  
	    lastPointer--  
	}    
}
```
- Go
```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
	lastPointer := m + n - 1  
  
	for m > 0 && n > 0 {  
	    if nums1[m-1] > nums2[n-1] {  
	       nums1[lastPointer] = nums1[m-1]  
	       m--  
	    } else {  
	       nums1[lastPointer] = nums2[n-1]  
	       n--  
	    }  
  
	    lastPointer--  
	}  
  
	for n > 0 {  
	    nums1[lastPointer] = nums2[n-1]  
	    n, lastPointer = n-1, lastPointer-1  
	}  
  
	fmt.Println("Nums: ", nums1)
}
```
- Dart
```dart
void merge(List<int> nums1, int m, List<int> nums2, int n) {
	var lastPointer = m + n - 1;  
	var pointerM = m - 1;  
	var pointerN = n - 1;  
  
	while (pointerM >= 0 && pointerN >= 0) {  
	  if (nums1[pointerM] > nums2[pointerN]) {  
	    nums1[lastPointer] = nums1[pointerM];  
	    pointerM--;  
	  } else {  
	    nums1[lastPointer] = nums2[pointerN];  
	    pointerN--;  
	  }  
	  
	  lastPointer--;  
	}  
  
	while (pointerN >= 0) {  
	  nums1[lastPointer] = nums2[pointerN];  
	  pointerN--;  
	  lastPointer--;  
	}
}
```
#### References
- [NeetCode](https://youtu.be/P1Ic85RarKY?si=azyJ9gGNlybT67t7)