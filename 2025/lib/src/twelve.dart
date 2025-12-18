import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';
import 'package:advent/src/math_utils.dart';

@Day(12)
class DayTwelveRunner implements DayRunner {
  @override
  String run() {
    Debugger.enable();
    var input = AdventInputReader("twelve.txt").readIntoLines();
    var fakePresentPacker = FakePresentPacker();
    var result = fakePresentPacker.getValidPresentRegions(input);
    return "$result";
  }
}

typedef Region = ({int area, List<int> shapeCounts});

class FakePresentPacker {
  int getValidPresentRegions(List<String> input) {
    var shapes = parseShapes(input);
    var regions = parseRegions(input);
    var validRegionCount = 0;
    for (var region in regions) {
      if (isRegionValid(region, shapes)) {
        validRegionCount++;
      }
    }
    return validRegionCount;
  }

  bool isRegionValid(Region region, List<int> shapes) {
    var totalPresentsArea = 0;
    for (var i = 0; i < region.shapeCounts.length; i++) {
      totalPresentsArea += region.shapeCounts[i] * shapes[i];
    }
    return totalPresentsArea <= region.area;
  }

  List<int> parseShapes(List<String> input) {
    var shapeIndexes = [1];
    var nextBlankIndex = input.indexWhere(
      (line) => line.trim().isEmpty,
      shapeIndexes.last,
    );
    while (nextBlankIndex != -1 && nextBlankIndex < 29) {
      shapeIndexes.add(nextBlankIndex + 2);
      nextBlankIndex = input.indexWhere(
        (line) => line.trim().isEmpty,
        shapeIndexes.last,
      );
    }

    var shapes = <int>[];

    for (var shapeIdx in shapeIndexes) {
      var area = 0;
      for (var i = shapeIdx; i < shapeIdx + 3; i++) {
        area += input[i]
            .split(".")
            .where((s) => s.trim().isNotEmpty)
            .join()
            .length;
      }
      shapes.add(area);
    }

    return shapes;
  }

  List<Region> parseRegions(List<String> input) {
    var regions = <Region>[];

    var startingIndex = input.lastIndexWhere((str) => str.trim().isEmpty) + 1;
    for (var i = startingIndex; i < input.length; i++) {
      var line = input[i];
      var split = line.split(":");
      var area = split[0].split("x").map(int.parse).reduce(product);
      var shapeCounts = split[1]
          .split(" ")
          .where((str) => str.trim().isNotEmpty)
          .map(int.parse)
          .toList();
      regions.add((area: area, shapeCounts: shapeCounts));
    }

    return regions;
  }
}
