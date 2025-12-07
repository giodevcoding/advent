import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';

@Day(5)
class DayFiveRunner implements DayRunner {
  @override
  String run() {
    var inputReader = AdventInputReader("five.txt");
    var input = inputReader.readIntoLines();
    var freshnessChecker = FreshnessChecker();
    var freshIngedientCount = freshnessChecker.getNumberOfFreshIngredients(
      input,
    );
    var totalPossibleCount = freshnessChecker
        .getTotalNumberofPossibleFreshIngredients(input);
    return "\nFresh Ingredient Count: $freshIngedientCount"
        "\nTotal Possible Fresh Ingredient Count: $totalPossibleCount";
  }
}

class FreshnessChecker {
  int getTotalNumberofPossibleFreshIngredients(List<String> rawInput) {
    var (ranges, _) = parseInput(rawInput);
    var normalizedRanges = normalizeRangeOverlaps(ranges);
    Debugger.log(normalizedRanges);
    Debugger.waitForInput();
    return normalizedRanges
        .map((range) => (range.high - range.low) + 1)
        .reduce((a, b) => a + b);
  }

  int getNumberOfFreshIngredients(List<String> rawInput) {
    var (freshRanges, availableIds) = parseInput(rawInput);
    var freshIds = getFreshIngredientIds(availableIds, freshRanges);
    return freshIds.length;
  }

  List<int> getFreshIngredientIds(
    List<int> availableIds,
    List<Range> freshRanges,
  ) {
    return availableIds
        .where((id) => isIngredientIdFresh(id, freshRanges))
        .toList();
  }

  bool isIngredientIdFresh(int ingredientId, List<Range> freshRanges) {
    for (var range in freshRanges) {
      if (ingredientId >= range.low && ingredientId <= range.high) {
        return true;
      }
    }
    return false;
  }

  (List<Range> freshRanges, List<int> availableIds) parseInput(
    List<String> rawInput,
  ) {
    var dividerIndex = rawInput.indexWhere((line) => line.trim().isEmpty);

    var rangeInput = rawInput.sublist(0, dividerIndex);
    var idInput = rawInput.sublist(dividerIndex + 1);

    var ranges = parseRanges(rangeInput);
    var ids = idInput.map(int.parse).toList();

    return (ranges, ids);
  }

  List<Range> parseRanges(List<String> rangeInput) => rangeInput.map((line) {
    var rawNumbers = line.split("-");
    return Range(int.parse(rawNumbers[0]), int.parse(rawNumbers[1]));
  }).toList();

  List<Range> normalizeRangeOverlaps(List<Range> startingRanges) {
    var oldRanges = List<Range>.from(startingRanges)..sort();
    var newRanges = <Range>[];

    for (var i = 0; i < oldRanges.length; i++) {
      var currentRange = oldRanges[i].copy();

      for (var j = i + 1; j < oldRanges.length; j++) {
        var nextRange = oldRanges[j];
        if (nextRange.low > currentRange.high) {
          break;
        }

        if (nextRange.high > currentRange.high) {
          currentRange.high = nextRange.high;
        }

        Debugger.log(currentRange);
        Debugger.waitForInput();
        i = j;
      }
      newRanges.add(currentRange);
    }

    return newRanges;
  }
}

class Range implements Comparable<Range> {
  int low, high;
  Range(this.low, this.high);

  Range copy() => Range(low, high);

  @override
  int compareTo(Range other) {
    var lowComparison = low.compareTo(other.low);
    if (lowComparison == 0) {
      return high.compareTo(other.high);
    }
    return lowComparison;
  }

  @override
  String toString() {
    return "{ low: $low, high: $high }";
  }
}
