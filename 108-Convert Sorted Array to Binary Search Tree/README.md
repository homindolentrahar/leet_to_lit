# 108. Convert Sorted Array to Binary Search Tree

#algorithm #data-structure #leetcode

#### Level

#leetcode-easy

#### Topics

#array #divide-conquer #tree #binary-search-tree #binary-tree

#### Description

Given an integer array`nums`where the elements are sorted in**ascending order**, convert_it to a_**_height-balanced_**
_binary search tree_.

#### Definition

Example 1:
![Example 1](https://github.com/homindolentrahar/leet_to_lit/blob/master/108-Convert%20Sorted%20Array%20to%20Binary%20Search%20Tree/assets/example_1.png)

```
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]

Explanation: [0,-10,5,null,-3,null,9] is also accepted
```

![Example 1](https://github.com/homindolentrahar/leet_to_lit/blob/master/108-Convert%20Sorted%20Array%20to%20Binary%20Search%20Tree/assets/example_1_result.png)

Example 2:
![Example 1](https://github.com/homindolentrahar/leet_to_lit/blob/master/108-Convert%20Sorted%20Array%20to%20Binary%20Search%20Tree/assets/example_2.png)

```
Input: nums = [1,3]
Output: [3,1]

Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.
```

Constraints:

- `1 <= nums.length <= 104`
- `-104 <= nums[i] <= 104`
- `nums`is sorted in a**strictly increasing**order.

#### Approaches

From the [[#Definition]] and [[#Description]], we know that:

1. `nums` is sorted in a **ASCENDING** order
2. We need to convert single array into Binary Search Tree
   Binary Search Tree will structure elements of `nums` based on their value in a `node` (the circle or dot). If other
   value is **LESSER** than a `node's` value, then it will assign to the left side (`left leaf` or  `left child`) of the
   node. Otherwise, it will assigned to the right side (`right leaf` or `right child`) of the node.
3. We will use [[Divide & Conquer Algorithm]] to breakdown `nums` into smaller element and then solve it independently
   This algorithm help us to solve a problem by breaking it down into smaller sub-problems, so that we can solve it
   independently, before merging the solutions to tackle the bigger or primary problem.
   We will also implement `recursive` approach by using this algorithm.

##### Divide & Conquer

1. Check if the `nums` is empty
   We want to catch the empty problem on the first line, so we won't deal with the empty array in the rest of function.
   If the `nums` is empty, means that we don't have any element to be converted into Binary Search Tree (BST for short),
   hence we cannot create a BST. We can return `null` or `nil` right away.
2. Define the `mid` pointer of `nums`
   We want to split `nums` into half on each `recursive` iteration, so we need a `mid` to determine where to split
   the `nums`. This is a part of the [[Divide & Conquer Algorithm]] that was mentioned before
3. Define the `node` object
   We need to define a `node` object that we will return later. `node` has a 3 properties:
    - Value
      Is current value of the `node`. We need assign it with `nums[mid]`, because we want the median of `nums` (or
      element at `mid` position) become the parent `node` of the both split `nums`.
    - Left Node
      Is a `left node` or `left children` of the current `node`. The left value usually is **LESSER** than
      current `node's` value.
      We need to assign it with `recursive` method, with the parameter of **LEFT SIDE** of the `nums` that has been
      split by `mid` pointer (sub-problem). We want to `recursively` run the function again to check and split the
      sub-problem into smaller sub-problem, so we can structure `nums` param into BST.
    - Right Node
      Is a `right node` or `right children` of the current `node`. The right value usually is **GREATER** than
      current `node's` value.
      We need to assign it with `recursive` method, with the parameter of **RIGHT SIDE** of the `nums` that has been
      split by `mid` pointer (sub-problem). We want to `recursively` run the function again to check and split the
      sub-problem into smaller sub-problem, so we can structure `nums` param into BST.
4. Return the `node`

#### Solution

- Kotlin

```kotlin
class TreeNode(var `val`: Int) {
    var left: TreeNode? = null
    var right: TreeNode? = null
}

fun sortedArrayToBST(nums: IntArray): TreeNode? {
    if (nums.isEmpty()) {
        return null
           
    }

        val mid = nums.size / 2
        val node = TreeNode(nums[mid])

        node.left = sortedArrayToBST(nums.sliceArray(0 until mid))
        node.right = sortedArrayToBST(nums.sliceArray(mid+1 until nums.size))

        return node
}
```

- Go

```go
type TreeNode struct {
Val   int
Left  *TreeNode
Right *TreeNode  
}

func sortedArrayToBST(nums []int) *TreeNode {
if len(nums) == 0 {
return nil
    }

    mid := len(nums) / 2
    node := &TreeNode{Val: nums[mid]}

    node.Left = sortedArrayToBST(nums[:mid])
    node.Right = sortedArrayToBST(nums[mid+1:])
   
    return node
}
```

- Dart

```dart
class TreeNode {
  int val;
  TreeNode? left;
  TreeNode? right;

  TreeNode([this.val = 0, this.left, this.right]);
}

TreeNode? sortedArrayToBST(List<int> nums) {
      if (nums.isEmpty) {
        return null;
      }

      final mid = nums.length ~/ 2;
      final node = TreeNode(nums[mid]);
     
      node.left = sortedArrayToBST(nums.sublist(0, mid));
      node.right = sortedArrayToBST(nums.sublist(mid + 1));
     
      return node;
  }
```

#### References

- [NeetCode](https://youtu.be/0K0uCMYq5ng?si=OH145shMvDll4gQn)
- [Divide & Conquer Algorithm in 3 Minutes](https://youtu.be/YOh6hBtX5l0?si=QznV76eRG0OZ8qEr)
- [Divide and Conquer: The Art of Breaking Down Problems | Recursion Series](https://youtu.be/ib4BHvr5-Ao?si=YbchQghH-mzVLM5p)