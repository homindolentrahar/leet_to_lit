main() {
  final nums = [-10, -3, 0, 5, 9];
  ConvertSortedArrayToBTS.sortedArrayToBST(nums);
}

class TreeNode {
  int val;
  TreeNode? left;
  TreeNode? right;

  TreeNode([this.val = 0, this.left, this.right]);
}

abstract class ConvertSortedArrayToBTS {
  static TreeNode? sortedArrayToBST(List<int> nums) {
    //
    // First attempt 07 June 2024 13:56
    // Runtime: 338ms,Memory: 148.70MB
    //
    if (nums.isEmpty) {
      return null;
    }

    final mid = nums.length ~/ 2;
    final node = TreeNode(nums[mid]);

    node.left = sortedArrayToBST(nums.sublist(0, mid));
    node.right = sortedArrayToBST(nums.sublist(mid + 1));

    return node;
  }
}
