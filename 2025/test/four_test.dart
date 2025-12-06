import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/four.dart';
import 'package:test/test.dart';

void main() {
  late ForkliftOptimizer forkliftOptimizer;
  late List<String> testInput;

  setUp(() {
    forkliftOptimizer = ForkliftOptimizer();
    testInput = [
      "..@@.@@@@.",
      "@@@.@.@.@@",
      "@@@@@.@.@@",
      "@.@@@@..@.",
      "@@.@@@@.@@",
      ".@@@@@@@.@",
      ".@.@.@.@@@",
      "@.@@@.@@@@",
      ".@@@@@@@@.",
      "@.@.@@@.@.",
    ];
  });

  test('ForkliftOptimizer.getAccessibleRolls returns correct number', () {
    var result = forkliftOptimizer.getAccessibleRolls(testInput).length;

    expect(result, 13);
  });

  test('ForkliftOptimizer.getTotalRemovableRollCount returns correct number', () {
    Debugger.enable();
    var result = forkliftOptimizer.getTotalRemovableRollCount(testInput);

    expect(result, 43);
  });

  tearDown(Debugger.disable);
}
