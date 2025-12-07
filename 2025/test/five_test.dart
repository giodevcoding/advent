import 'dart:developer';

import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/five.dart';
import 'package:test/test.dart';

void main() {
  late List<String> testInput;
  late FreshnessChecker freshnessChecker;

  setUp(() {
    testInput = [
      "3-5",
      "10-14",
      "12-18",
      "12-18",
      "12-19",
      "13-17",
      "16-20",
      "",
      "1",
      "5",
      "8",
      "11",
      "17",
      "32",
    ];
    freshnessChecker = FreshnessChecker();
  });
  test(
    'FreshnessChecker.getNumberOfFreshIngredients returns correct number',
    () {
      var result = freshnessChecker.getNumberOfFreshIngredients(testInput);

      expect(result, 3);
    },
  );

  test(
    'FreshnessChecker.getTotalNumberofPossibleFreshIngredients returns correct number',
    () {
      var result = freshnessChecker.getTotalNumberofPossibleFreshIngredients(testInput);

      expect(result, 14);
    },
  );

  tearDown(Debugger.disable);
}
