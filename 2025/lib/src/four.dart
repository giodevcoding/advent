import 'dart:math';

import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';
import 'package:advent/src/string_utils.dart';

@Day(4)
class DayFourRunner implements DayRunner {
  @override
  String run() {
    var fileReader = AdventInputReader("four.txt");
    var warehouseGrid = fileReader.readIntoLines();
    var forkliftOptimizer = ForkliftOptimizer();
    var initialAccessibleRollCount = forkliftOptimizer
        .getAccessibleRolls(warehouseGrid)
        .length;
    var totalRemovableRollCount = forkliftOptimizer.getTotalRemovableRollCount(
      warehouseGrid,
    );
    return "\nInitial Accessible Rolls: $initialAccessibleRollCount\n"
        "Total Removable Rolls: $totalRemovableRollCount";
  }
}

class ForkliftOptimizer {
  int getTotalRemovableRollCount(List<String> warehouseGrid, {int count = 0}) {
    var accessibleRolls = getAccessibleRolls(warehouseGrid);
    if (accessibleRolls.isEmpty) {
      return count;
    }

    _debugLogGridBeforeRemoval(warehouseGrid, accessibleRolls);
    var gridAfterRemoval = _getGridAfterRemoval(warehouseGrid, accessibleRolls);
    _debugLogGrid(gridAfterRemoval);
    return getTotalRemovableRollCount(
      gridAfterRemoval,
      count: count + accessibleRolls.length,
    );
  }

  List<Tile> getAccessibleRolls(List<String> warehouseGrid) {
    var accessibleRolls = <Tile>[];
    _scanGrid(warehouseGrid, (x, y) {
      var tile = _getTileAt(x, y, warehouseGrid)!;

      if (tile.tileValue != TileValue.paperRoll) {
        return;
      }

      if (_isRollTileAccessible(tile, warehouseGrid)) {
        accessibleRolls.add(tile);
      }
    });

    return accessibleRolls;
  }

  void _scanGrid(List<String> grid, Function(int x, int y) scanCallback) {
    for (var y = 0; y < grid.length; y++) {
      for (var x = 0; x < grid[0].length; x++) {
        scanCallback(x, y);
      }
    }
  }

  bool _isRollTileAccessible(Tile tile, List<String> warehouseGrid) {
    var surroundingTiles = _getSurroundingTiles(tile, warehouseGrid);
    var rollCount = _getRollsCount(surroundingTiles);
    return rollCount < 4;
  }

  int _getRollsCount(List<Tile> tiles) =>
      tiles.where((t) => t.tileValue == TileValue.paperRoll).length;

  List<Tile> _getSurroundingTiles(Tile tile, List<String> warehouseGrid) {
    var tiles = <Tile?>[];
    for (var yOffset = -1; yOffset <= 1; yOffset++) {
      for (var xOffset = -1; xOffset <= 1; xOffset++) {
        if (yOffset == 0 && xOffset == 0) {
          continue;
        }

        var xCheck = tile.x + xOffset;
        var yCheck = tile.y + yOffset;
        tiles.add(_getTileAt(xCheck, yCheck, warehouseGrid));
      }
    }

    return tiles.nonNulls.toList();
  }

  Tile? _getTileAt(int x, int y, List<String> warehouseGrid) {
    var yOutOfBounds = y < 0 || y >= warehouseGrid.length;
    var xOutOfBounds = x < 0 || x >= warehouseGrid[0].length;

    if (yOutOfBounds || xOutOfBounds) {
      return null;
    }

    var point = Point(x, y);
    var value = TileValue.fromValue(warehouseGrid[y][x]);
    return Tile(value, point);
  }

  List<String> _getGridAfterRemoval(
    List<String> warehouseGrid,
    List<Tile> accessibleTiles,
  ) {
    var newGrid = List<String>.from(warehouseGrid);
    for (var tile in accessibleTiles) {
      newGrid[tile.y] = replaceCharAt(newGrid[tile.y], tile.x, ".");
    }

    return newGrid;
  }

  void _debugLogGrid(List<String> warehouseGrid) {
    warehouseGrid.forEach(Debugger.log);
    Debugger.waitForInput();
  }

  void _debugLogGridBeforeRemoval(
    List<String> warehouseGrid,
    List<Tile> accessibleRolls,
  ) {
    Debugger.log("Removing ${accessibleRolls.length} rolls");
    var logGrid = List<String>.from(warehouseGrid);
    for (var tile in accessibleRolls) {
      logGrid[tile.y] = replaceCharAt(logGrid[tile.y], tile.x, "x");
    }

    logGrid.forEach(Debugger.log);
    Debugger.waitForInput();
  }
}

class Tile {
  final TileValue tileValue;
  final Point position;

  const Tile(this.tileValue, this.position);

  @override
  String toString() {
    return "{value: $value, x: $x, y: $y}";
  }

  int get x => position.x.floor();
  int get y => position.y.floor();
  String get value => tileValue.value;
}

enum TileValue {
  paperRoll('@'),
  empty('.');

  final String value;

  const TileValue(this.value);

  static TileValue fromValue(String val) =>
      TileValue.values.firstWhere((tile) => tile.value == val);
}
