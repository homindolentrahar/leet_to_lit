main() {
  final nums = [2, 2, 1];
  SingleNumber.singleNumber(nums);
}

abstract class SingleNumber {
  static int singleNumber(List<int> nums) {
    //
    // First attempt 25 June 2024 16:06
    // Runtime: 346ms, Memory: 150.20MB
    //
    int val = 0;

    for (int i = 0; i < nums.length; i++) {
      val = val ^ nums[i];
    }

    return val;
  }
}
