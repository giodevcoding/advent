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

  test('TachyonBeamTracker.splitCount returns correct number', () {
    tachyonBeamTracker = TachyonBeamTracker(testInput);

    expect(tachyonBeamTracker.splitCount, 21);
  });

  test('TachyonBeamTracker.pathCount returns correct number', () {
    Debugger.enable();
    tachyonBeamTracker = TachyonBeamTracker(testInput);
    var pathCount = tachyonBeamTracker.getTimelineCount();

    expect(pathCount, 40);
  });

  tearDown(Debugger.disable);
}
