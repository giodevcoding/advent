import 'dart:math';
import 'package:collection/collection.dart';

import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';
import 'package:advent/src/math_utils.dart';

@Day(9)
class DayNineRunner implements DayRunner {
  @override
  String run() {
    var inputReader = AdventInputReader("nine.txt");
    var input = inputReader.readIntoLines();
    var redTheaterTileAnalyzer = RedTheaterTileAnalyzer();
    var largestRectangleRed = redTheaterTileAnalyzer.getLargestRectangle(input);
    var redGreenTheaterTileAnalyzer = RedGreenTheaterTileAnalyzer();
    var largestRectangleRedGreen = redGreenTheaterTileAnalyzer
        .getLargestRectangle(input);

    return "\nLargest Rectangle with Red Tiles: $largestRectangleRed"
        "\nLargest Rectangle with Red+Green Tiles: $largestRectangleRedGreen";
  }
}

abstract class AbstractTheaterTileAnalyzer {
  int getLargestRectangle(List<String> input);

  List<TheaterTile> getRedTiles(List<String> input) => input
      .map((line) => line.split(",").map(int.parse).toList())
      .map((coords) => TheaterTile(coords[0], coords[1]))
      .toList();
}

class RedTheaterTileAnalyzer extends AbstractTheaterTileAnalyzer {
  @override
  int getLargestRectangle(List<String> input) {
    var redTiles = getRedTiles(input);
    return _getLargestRectangleFromRedTiles(redTiles);
  }

  int _getLargestRectangleFromRedTiles(List<TheaterTile> redTiles) {
    var largest = 0;

    for (var i = 0; i < redTiles.length - 1; i++) {
      for (var j = i + 1; j < redTiles.length; j++) {
        var rectangle = TheaterRectangle(redTiles[i], redTiles[j]);
        largest = max(rectangle.area, largest);
      }
    }

    return largest;
  }
}

class RedGreenTheaterTileAnalyzer extends AbstractTheaterTileAnalyzer {
  @override
  int getLargestRectangle(List<String> input) {
    Debugger.enable();
    var redTiles = getRedTiles(input);
    var rectangles = _getRectangles(redTiles);
    var verticalEdges = _getVerticalEdges(redTiles);
    var horizontalRanges = _getHorizontalRanges(verticalEdges);

    for (var i = 0; i < rectangles.length; i++) {
      var rect = rectangles[i];
      Debugger.log("Checking rectangle $i of ${rectangles.length}");
      if (_isRectangleValid(rect, verticalEdges, horizontalRanges)) {
        return rect.area;
      }
    }

    return 0;
  }

  List<TheaterRectangle> _getRectangles(List<TheaterTile> redTiles) {
    var rectangles = <TheaterRectangle>[];
    for (var i = 0; i < redTiles.length - 1; i++) {
      for (var j = i + 1; j < redTiles.length; j++) {
        rectangles.add(TheaterRectangle(redTiles[i], redTiles[j]));
      }
    }

    return rectangles.sorted((r1, r2) => r2.area.compareTo(r1.area)).toList();
  }

  bool _isRectangleValid(
    TheaterRectangle rectangle,
    List<VerticalEdge> verticalEdges,
    Map<int, List<HorizontalRange>> horizontalRanges,
  ) {
    for (var y = rectangle.minY; y <= rectangle.maxY; y++) {
      var xIncrement = y == rectangle.minY || y == rectangle.maxY ? 1 : rectangle.maxX - rectangle.minX;
      for (var x = rectangle.minX; x <= rectangle.maxX; x += xIncrement) {
        var inRange = false;
        for (var range in horizontalRanges[y]!) {
          if (x >= range.minX && x <= range.maxX) {
            inRange = true;
          }
        }

        if (!inRange) {
          return false;
        }
      }
    }
    return true;
  }

  List<VerticalEdge> _getVerticalEdges(List<TheaterTile> redTiles) {
    var verticalEdges = <VerticalEdge>[];

    for (var i = 0; i < redTiles.length; i++) {
      var nextTileIndex = i + 1;
      if (i == redTiles.length - 1) {
        nextTileIndex = 0;
      }

      var currentRedTile = redTiles[i];
      var nextRedTile = redTiles[nextTileIndex];

      if (currentRedTile.y == nextRedTile.y) {
        continue;
      }

      if (currentRedTile.x != nextRedTile.x) {
        continue;
      }

      VerticalEdge edge = (
        x: currentRedTile.x,
        minY: min(currentRedTile.y, nextRedTile.y),
        maxY: max(currentRedTile.y, nextRedTile.y),
      );
      verticalEdges.add(edge);
    }

    return verticalEdges;
  }

  Map<int, List<HorizontalRange>> _getHorizontalRanges(
    List<VerticalEdge> verticalEdges,
  ) {
    var horizontalRanges = <int, List<HorizontalRange>>{};

    int minY = double.maxFinite.toInt(), maxY = 0;
    for (var edge in verticalEdges) {
      minY = min(minY, edge.minY);
      maxY = max(maxY, edge.maxY);
    }

    var edges = verticalEdges.sorted((e1, e2) => e1.x.compareTo(e2.x));

    for (var y = minY; y <= maxY; y++) {
      var ranges = <HorizontalRange>[];
      VerticalEdge? startingEdge;

      for (var edge in edges) {
        var minY = edge.minY;
        var maxY = edge.maxY;
        if (startingEdge != null) {
          if (y == startingEdge.minY) {
            maxY--;
          }
          if (y == startingEdge.maxY) {
            minY++;
          }
        }

        if (y >= minY && y <= maxY) {
          if (startingEdge == null) {
            startingEdge = edge;
          } else {
            ranges.add((minX: startingEdge.x, maxX: edge.x, y: y));
            startingEdge = null;
          }
        }
      }
      horizontalRanges[y] = ranges;
    }

    return horizontalRanges;
  }
}

class TheaterTile {
  final int x, y;

  const TheaterTile(this.x, this.y);

  @override
  bool operator ==(Object other) =>
      other is TheaterTile &&
      other.runtimeType == runtimeType &&
      other.x == x &&
      other.y == y;

  @override
  int get hashCode => x.hashCode * y.hashCode;

  @override
  String toString() => "{$x,$y}";
}

class TheaterRectangle {
  final TheaterTile tileA, tileB;
  int? _area;

  TheaterRectangle(this.tileA, this.tileB);

  int get area {
    if (_area == null) {
      var xDistance = (tileB.x - tileA.x).abs() + 1;
      var yDistance = (tileB.y - tileA.y).abs() + 1;
      _area = xDistance * yDistance;
    }

    return _area!;
  }

  int get minX => min(tileA.x, tileB.x);
  int get minY => min(tileA.y, tileB.y);
  int get maxX => max(tileA.x, tileB.x);
  int get maxY => max(tileA.y, tileB.y);

  @override
  String toString() => "{tileA: $tileA, tileB: $tileB, area: $area}";
}

typedef VerticalEdge = ({int x, int minY, int maxY});
typedef HorizontalRange = ({int y, int minX, int maxX});
