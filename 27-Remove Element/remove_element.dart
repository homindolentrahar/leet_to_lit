main() {
  final nums = [3, 2, 2, 3];
  final val = 3;

  final result = RemoveElement.removeElement(nums, val);
  print(result);
}

abstract class RemoveElement {
  static int removeElement(List<int> nums, int val) {
    //
    // First attempt May 31, 2024 21:53
    // Runtime 337ms, Memory 151.24MB
    //
    // int pointer = 0;
    //
    // for (var index = 0; index < nums.length; index++) {
    //   if (nums[index] != val) {
    //     nums[pointer] = nums[index];
    //     pointer++;
    //   }
    // }
    //
    // return pointer;
    //
    // Second attempt May 31, 2024 21:56
    // Runtime 301ms, Memory 147.06MB
    //
    for (int i = 0; i < nums.length; i++) {
      if (nums[i] == val) {
        nums.removeAt(i);
        i--;
      }
    }

    return nums.length;
  }
}
