main() {
  final rowIndex = 4;
  PascalTriangle2.getRow(rowIndex);
}

abstract class PascalTriangle2 {
  static List<int> getRow(int rowIndex) {
    //
    // First attempt 26 June 2024 13:09
    // Runtime: 330ms, Memory: 146.70MB
    //
    // final res = [1];
    // var prev = 1;
    //
    // for (int i = 1; i <= rowIndex; i++) {
    //   final curr = prev.toDouble() * (rowIndex - i + 1) ~/ i;
    //   res.add(curr);
    //   prev = curr;
    // }
    //
    // return res;
    //
    // Second attempt 26 June 2024
    // Runtime: 300ms, Memory: 151.76MB
    // Recursive
    //
    if (rowIndex == 0) return [1];

    final prevRow = getRow(rowIndex - 1);
    final row = <int>[];

    row.add(1);
    for (int i = 1; i < rowIndex; i++) {
      row.add(prevRow[i - 1] + prevRow[i]);
    }
    row.add(1);

    return row;
  }
}
