main() {
  final num = 5;
  PascalTriangle.generate(num);
}

abstract class PascalTriangle {
  static List<List<int>> generate(int numRows) {
    //
    // First attempt 25 June 2024 12:49
    // Runtime: 308ms, Memory: 146.33MB
    // For loop
    //
    // final res = [
    //   [1]
    // ];
    //
    // for (int i = 0; i < numRows - 1; i++) {
    //   final lastRes = res[res.length - 1];
    //   final temp = [0, ...lastRes, 0];
    //   final row = <int>[];
    //
    //   for (int j = 0; j < lastRes.length + 1; j++) {
    //     row.add(temp[j] + temp[j + 1]);
    //   }
    //
    //   res.add(row);
    // }
    //
    // return res;
    //
    // Second attempt 25 June 2024 12:54
    // Runtime: 323ms, Memory: 145.22MB
    // Combinatorial Factorial
    //
    // final res = [[1]];
    //
    // for (int i = 1; i < numRows; i++) {
    //   final prevRows = res[i - 1];
    //   final row = <int>[];
    //
    //   row.add(1);
    //   for (int j = 1; j < i; j++) {
    //     row.add(prevRows[j - 1] + prevRows[j]);
    //   }
    //   row.add(1);
    //   res.add(row);
    // }
    //
    // return res;
    //
    // Third attempt 25 June 2024 12:59
    // Runtime: 332ms, Memory: 145.12MB
    // Dynamic Programming with 1D Array
    //
    if (numRows == 0) {
      return [[]];
    }

    if (numRows == 1) {
      return [[1]];
    }

    final res = List.of(generate(numRows - 1));
    final prevRow = res[numRows - 2];
    final row = <int>[];

    row.add(1);
    for (int i = 1; i < numRows - 1; i++) {
      row.add(prevRow[i - 1] + prevRow[i]);
    }
    row.add(1);
    res.add(row);

    return res;
  }
}
