import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/nine.dart';
import 'package:test/test.dart';

void main() {
  late RedTheaterTileAnalyzer redTheaterTileAnalyzer;
  late RedGreenTheaterTileAnalyzer redGreenTheaterTileAnalyzer;
  late List<String> testInput;

  setUp(() {
    redTheaterTileAnalyzer = RedTheaterTileAnalyzer();
    redGreenTheaterTileAnalyzer = RedGreenTheaterTileAnalyzer();
    testInput = ["7,1", "11,1", "11,7", "9,7", "9,5", "2,5", "2,3", "7,3"];
  });

  test('RedTheaterTileAnalyzer.getLargestRectangle returns expected size', () {
    var actual = redTheaterTileAnalyzer.getLargestRectangle(testInput);

    expect(actual, 50);
  });

  test('RedGreenTheaterTileAnalyzer.getLargestRectangle returns expected size', () {
    Debugger.enable();
    var actual = redGreenTheaterTileAnalyzer.getLargestRectangle(testInput);

    expect(actual, 24);
  });

  tearDown(Debugger.disable);
}
