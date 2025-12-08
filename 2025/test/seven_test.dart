import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/seven.dart';
import 'package:test/test.dart';

void main() {
  late List<String> testInput;
  late TachyonBeamTracker tachyonBeamTracker;
  setUp(() {
    testInput = [
      ".......S.......",
      "...............",
      ".......^.......",
      "...............",
      "......^.^......",
      "...............",
      ".....^.^.^.....",
      "...............",
      "....^.^...^....",
      "...............",
      "...^.^...^.^...",
      "...............",
      "..^...^.....^..",
      "...............",
      ".^.^.^.^.^...^.",
      "...............",
    ];
  });

  test('TachyonBeamTracker.getSplitCount returns correct number', () {
    Debugger.enable();
    tachyonBeamTracker = TachyonBeamTracker(testInput);

    expect(tachyonBeamTracker.splitCount, 21);
  });

  tearDown(Debugger.disable);
}
