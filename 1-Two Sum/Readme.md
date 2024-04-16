# 1. Two Sums
#data-structure #algorithm #leetcode
#### Level
#leetcode-easy 
#### Topics
#array #hash-table
#### Description
Given an array of integers `nums` and an integer `target`, return _indices of the two numbers such that they add up to `target`_. Return the index of two values which equal to target if both are added. 
> ❌ Cannot use the same element twice
#### Definition
Example 1:
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
```
Example 2:
```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```
Example 3:
```
Input: nums = [3,3], target = 6
Output: [0,1]
```
Constrains:
- `2 <= nums.length <= 104`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
#### Approaches
##### Brute Force - O(n<sup>2</sup>) (❌ Very Slow)
1. Create first `for` loop iterating the `nums` variable (Save the index as `i`)
2. For every iteration, create another `for` loop (inner loop) that runs from index `i + 1` < `nums.size` 
3. Check, if `nums` with outer index (`nums[i]`) is equal to `nums` with inner index (`nums[j]`). If `true`, then assign `indexes` (which holds the array with indices index) with outer and inner index (`i` and `j`)
4. Finally, return the `indexes`
```kotlin
var indexes = intArrayOf()  
for (i in nums.indices) {  
    for (j in i + 1..<nums.size) {  
        if (nums[i] + nums[j] == target) {  
            indexes = intArrayOf(i, j)  
        }  
    }
}  
```
##### HashMap 
The formula of this equation can be down to as simple as
> `x + y = target`

 We already knew that:
- `target` is already pre-defined (e.g **9**)
- One of the variable (either `x` or `y`) is element of `nums`
- You cannot add the same element twice

So based on that information, we can tweak the formula into
> `y = target - x `

to find off the remaining value that add up to `target`. So we will utilize `HashMap` to store the `y` with the current index, then check if `x` is present at the `HashMap`. Now why'd we do that?

Because we don't have to find all the occurrences of all number. We just need two numbers that add up to target. Let say the target is `9` and the current element of `nums` is 7, so
> y = 9 - 7
> y = 2

We'll store `2` inside `HashMap` with the index. Then we need to check `x`, which `7`; is present in the `HashMap`. If `7` is present, it means that we already do the 
>y = 9 - 2
>y = 7

things, right? Since those two are interchangeable, we can take the current index of `7` and the index stored with `2` key from `HashMap`.

We can do this by:
1. Create a `HashMap` to holds `y` and it's index
2. Create an `Array` to holds the indices indexes
3. Find the `diff` or `y` with `y = target - x`; `x` represent current element of `nums`
4. Check if `x` is present in the `HashMap`
	- If present, then assign the `Array` with the index of `x` and the index that saved with key of `y`  from `HashMap`
	- If not, assign value in `HashMap` with key of `y` and value of `y`'s index
5. Return the `Array`
#### Solution
- Kotlin
```kotlin
fun twoSum(nums: IntArray, target: Int): IntArray {
	val hashTemp = hashMapOf<Int, Int>()
	var temp = intArrayOf()
	
	for (i in nums.indices) {
		val diffs = target - nums[i]
		
		if (hashTemp.keys.contains(diffs)) {
			temp = intArrayOf(hashTemp[diffs] ?: 0, i)
		}
		
		hashTemp[nums[i]] = i
	}  
  
	return temp
}
```
- Go
```go
func twoSum(nums []int, target int) []int {  
    for index, num := range nums {  
       if _, ok := mapResult[num]; ok {  
          return []int{mapResult[num], index}  
       }  
       diff := target - num  
       mapResult[diff] = index  
    }  
  
    return []int{}
}
```
#### References
- [AlgoEngine](https://youtu.be/luicuNOBTAI)