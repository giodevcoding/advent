import 'dart:math';

import 'package:advent/day.dart';
import 'package:advent/src/file_utils.dart';
import 'package:advent/src/debug_utils.dart';

@Day(1)
class DayOneRunner implements DayRunner {
  @override
  String run() {
    var fileReader = AdventInputReader("one.txt");
    var instructions = fileReader.readIntoLines();
    var decoder = NorthPoleEntrancePasswordDecoder(startingPoint: 50);
    var decodedPassword = decoder.decode(instructions);
    var decodedSpecialPassword = decoder.decodeSpecial(instructions);
    return "Decoded with Old Method: $decodedPassword\nDecode with method 0x434C49434B: $decodedSpecialPassword";
  }
}

class NorthPoleEntrancePasswordDecoder {
  final int startingPoint;

  NorthPoleEntrancePasswordDecoder({required this.startingPoint});

  int decode(List<String> instructions) {
    var currentDialPoint = startingPoint;
    var zeroHits = 0;

    for (var instruction in instructions) {
      var (direction, distance) = parseInstruction(instruction);

      (currentDialPoint, _) = rotateDialWithPasses(
        currentDialPoint,
        direction,
        distance,
      );

      if (currentDialPoint == 0) {
        zeroHits++;
      }
    }

    return zeroHits;
  }

  int decodeSpecial(List<String> instructions) {
    var currentDialPoint = startingPoint;
    var zeroPasses = 0;

    for (var instruction in instructions) {
      var (direction, distance) = parseInstruction(instruction);
      var passes = 0;

      Debugger.log(
        "$currentDialPoint ${direction == Direction.left ? '-' : '+'} $distance",
      );
      (currentDialPoint, passes) = rotateDialWithPasses(
        currentDialPoint,
        direction,
        distance,
      );
      Debugger.log("   + $passes passes");

      zeroPasses += passes;

      if (currentDialPoint == 0) {
        zeroPasses++;
      }

      Debugger.log("   = $currentDialPoint w/ $zeroPasses total zero hits");
      Debugger.waitForInput(when: zeroPasses > 39);
    }

    return zeroPasses;
  }

  (int, int) rotateDialWithPasses(
    int currentDialPoint,
    Direction direction,
    int distance,
  ) {
    var newDialPoint = currentDialPoint + (distance * direction.multiplier);

    var overlaps = ((newDialPoint / 100).floor() * 100 * direction.multiplier);

    if (newDialPoint < 0) {
      newDialPoint += overlaps;
    } else if (newDialPoint > 99) {
      newDialPoint -= overlaps;
    }

    var overlapsCount = (overlaps / 100).abs().round();
    Debugger.log("overlaps: $overlaps, count: $overlapsCount");
    if ((currentDialPoint == 0 && direction == Direction.left) ||
        (newDialPoint == 0 && direction == Direction.right)) {
      overlapsCount = max(overlapsCount - 1, 0);
    }

    return (newDialPoint, overlapsCount);
  }

  (Direction direction, int distance) parseInstruction(String instruction) {
    var direction = Direction.fromString(instruction.substring(0, 1));
    var distance = int.parse(instruction.substring(1));
    return (direction, distance);
  }
}

enum Direction {
  left(-1),
  right(1);

  final int multiplier;

  const Direction(this.multiplier);

  static Direction fromString(String value) {
    return switch (value.toUpperCase()) {
      "L" || "LEFT" => left,
      "R" || "RIGHT" => right,
      _ => throw ArgumentError(
        "Not a valid argument to parse a direction from",
      ),
    };
  }
}
