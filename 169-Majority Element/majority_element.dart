main() {
  final nums = [2, 3, 2];
  MajorityNumber.majorityElement(nums);
}

abstract class MajorityNumber {
  static int majorityElement(List<int> nums) {
    //
    // First attempt 01 July 2024 16:57
    // Runtime: 447ms, Memory: 152.79MB
    // Time complexity: O(n), Space complexity: O(1)
    // Boyer-Moore Voting Algorithm
    //
    int candidate = 0;
    int count = 0;

    for (int i = 0; i < nums.length; i++) {
      if (count == 0) {
        candidate = nums[i];
      }

      if (candidate == nums[i]) {
        count++;
      } else {
        count--;
      }
    }

    return candidate;
  }
}
