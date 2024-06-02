main() {
  final nums = [4, 3, 2, 1];

  final result = PlusOne.plusOne(nums);
  print(result);
}

abstract class PlusOne {
  static List<int> plusOne(List<int> digits) {
    //
    // First attempt Jun 02 2024 13:53
    // Runtime: 340ms, Memory: 148.24MB
    //
    for (int i = digits.length - 1; i >= 0; i--) {
      if (digits[i] != 9) {
        digits[i] = digits[i] + 1;
        return digits;
      }

      digits[i] = 0;
    }

    digits.insert(0, 1);
    return digits;
  }
}
