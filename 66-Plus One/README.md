# 66. Plus One
#data-structure #algorithm #leetcode 
#### Level
#leetcode-easy 
#### Topics
#array #math 
#### Description
You are given a **large integer** represented as an integer array `digits`, where each `digits[i]` is the `ith` digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading `0`'s.
Increment the large integer by one and return _the resulting array of digits_.
#### Definition
Example 1:
```
Input: digits = [1,2,3]
Output: [1,2,4]

Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
```
Example 2:
```
Input: digits = [4,3,2,1]
Output: [4,3,2,2]

Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
```
Example 3:
```
Input: digits = [9]
Output: [1,0]

Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].
```
Constraints:
-  `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`
- `digits` does not contain any leading `0`'s.
#### Approaches
From the [[#Definition]] and [[#Description]], we know that:
1. `digits` sorted from most significant to least significant in left-to-right order
2. We only want to increment the last element of `digits`
3. There are **NO** leading `0` in the `digits`
4. If the element of `digits` is `9`, it'll turn the element into `0` while carrying `1`
5. We will use `loop` approach, because it's basically not a single addition operation
	Imagine if the `digits` is `[9, 9, 9]`. That means you're not only perform operation on the last digit, because it will gives you `[9, 9, 1, 0]`. Instead, we would perform operation by loop that iterate `digits` from the **BACKWARDS** until it reaches the front (Because in the real world, that's how addition works).
##### Reversed Loop
1. Define the reversed loop
	Starts from the last index to the first index
2. Compare the current element with `9`
	- If the element is **NOT EQUAL** to `9`, then increment element by `1` and return the `digits` right away
		If the element is **NOT EQUAL** to `9`, that means we don't have to traverse all the way to the first index. We just need to increment last digit of `digits` by `1`, okay? So we stop the loop and return the incremented `digits` right away to prevent further iteration.
	- If the element is **EQUAL** to `9`, then replace the current element with `0`
		If you increment `9` by `1`, it will gives you `10`. But if the `digits` looks like `[2, 9, 9]`, it will carry the `1` up to the more significant one, hence we got `[3, 0, 0]`. 
		"*Then why we just need to replace it with `0`, didn't we have to handle the carry also?*" It will handled by the first case (first `if` checking). So if the more significant digit is **NOT EQUAL** to `9`, it will increment it by `1`. Otherwise, it will replace the digit with `0` and the loop continue
3. Return the `digits`, by append `1` on the `head`
	"*Why'd we need to do that? Doesn't it turn the `digits` from `[4, 3, 2, 1]` into `[1, 4, 3, 2, 1]`*" . No it doesn't. Because the moment loop hits the digit that **NOT EQUAL** to `9`, it will increment the digit by `1` and return the `digits` right away, so the `return` code down below doesn't get executed. Same thing if the `digits` is `[2, 9, 9]`, it will handled by the first and second case, the `return` code down below also doesn't get executed.
	We need this to handle edge-cases, if we encounter `digits` that has all elements in `9`, for example `[9, 9, 9, 9]`.  If we don't append `1` on the `head`, it will return `[0, 0, 0, 0]` (because it's only handled by the second case) which doesn't make sense in addition operation.
#### Solution
- Kotlin
```kotlin
fun plusOne(digits: IntArray): IntArray {  
	for (index in (digits.size - 1) downTo 0) {  
        if (digits[index] < 9) {  
            digits[index]++  
            return digits  
        }  
  
        digits[index] = 0  
    }  
  
    val newDigits = IntArray(digits.size + 1)  
    newDigits[0] = 1  
  
    return newDigits  
}
```
- Go
```go
func plusOne(digits []int) []int {  
    for i := len(digits) - 1; i >= 0; i-- {  
       if digits[i] != 9 {  
          digits[i]++  
          return digits  
       }  
  
       digits[i] = 0  
    }  
  
    return append([]int{1}, digits...)  
}
```
- Dart
```dart
List<int> plusOne(List<int> digits) {
    for(int i=digits.length-1;i>=0;i--){
        if(digits[i]<9){
            digits[i]++;
            return digits;
        }
        digits[i] = 0;
    }
    digits.insert(0,1);
    
	return digits;
}
```
#### References
- [AlgoEngine](https://youtu.be/RQAQA1FH_F4?si=nQXHYsVlN9s2YMZR)
- [NeetCode](https://youtu.be/jIaA8boiG1s?si=adRoHjw8kaYXy57w)