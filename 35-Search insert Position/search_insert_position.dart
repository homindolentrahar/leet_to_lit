main() {
  final nums = [3, 2, 2, 3];
  final val = 3;

  final result = SearchInsertPosition.searchInsert(nums, val);
  print(result);
}

abstract class SearchInsertPosition {
  static int searchInsert(List<int> nums, int target) {
    //
    // First attempt Jun 01 2024 11:51
    // Runtime: 331ms, Memory: 150.54MB
    //
    int start = 0;
    int end = nums.length - 1;

    while (start <= end) {
      final med = start + (end - start) ~/ 2;

      if (target == nums[med]) {
        return med;
      }

      if (target > nums[med]) {
        start = med + 1;
      } else {
        end = med - 1;
      }
    }

    return start;
  }
}
