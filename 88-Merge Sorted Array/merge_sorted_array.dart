main() {
  final nums1 = [1, 2, 3, 0, 0, 0];
  final nums2 = [2, 5, 6];
  final m = 3;
  final n = 3;

  MergeSortedArray.merge(nums1, m, nums2, n);
}

abstract class MergeSortedArray {
  static void merge(List<int> nums1, int m, List<int> nums2, int n) {
    //
    // First attempt 06 June 2024 15:38
    // Runtime: 367ms, Memory: 147.54MB
    //
    var lastPointer = m + n - 1;
    var pointerM = m - 1;
    var pointerN = n - 1;

    while (pointerM >= 0 && pointerN >= 0) {
      if (nums1[pointerM] > nums2[pointerN]) {
        nums1[lastPointer] = nums1[pointerM];
        pointerM--;
      } else {
        nums1[lastPointer] = nums2[pointerN];
        pointerN--;
      }

      lastPointer--;

    }

    while (pointerN >= 0) {
      nums1[lastPointer] = nums2[pointerN];
      pointerN--;
      lastPointer--;
    }
  }
}
