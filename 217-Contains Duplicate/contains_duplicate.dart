main() {
  final nums = [2, 3, 2];
  ContainsDuplicate.containsDuplicate(nums);
}

abstract class ContainsDuplicate {
  static bool containsDuplicate(List<int> nums) {
    //
    // First attempt 02 July 2024 11:09
    // Runtime: 400ms, Memory: 180.35MB
    // Time complexity: O(n), Space complexity: O(n)
    // HashSet
    //
    final occurrenceMap = <int>{};

    for (int i = 0; i < nums.length; i++) {
      if (occurrenceMap.contains(nums[i])) {
        return true;
      }

      occurrenceMap.add(nums[i]);
    }

    return false;
  }
}
