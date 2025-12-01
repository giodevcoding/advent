import 'package:advent/src/file_utils.dart';

String dayOne() {
  var fileReader = AdventInputReader("one.txt");
  var instructions = fileReader.readIntoLines();
  var decoder = NorthPoleEntrancePasswordDecoder(startingPoint: 50);
  var decodedPassword = decoder.decode(instructions);
  var decodedSpecialPassword = decoder.decodeSpecial(instructions);
  return "Decoded with Old Method: $decodedPassword\nDecode with method 0x434C49434B: $decodedSpecialPassword";
}

class NorthPoleEntrancePasswordDecoder {
  final int startingPoint;

  NorthPoleEntrancePasswordDecoder({required this.startingPoint});

  int decode(List<String> instructions) {
    var currentDialPoint = startingPoint;
    var zeroHits = 0;

    for (var instruction in instructions) {
      var (direction, distance) = parseInstruction(instruction);

      currentDialPoint = rotateDial(currentDialPoint, direction, distance);

      if (currentDialPoint == 0) {
        zeroHits++;
      }
    }

    return zeroHits;
  }

  int decodeSpecial(List<String> instructions) {
    return 0;
  }

  int rotateDial(int currentDialPoint, Direction direction, int distance) {
    if (direction == Direction.left) {
      currentDialPoint -= distance;
    } else if (direction == Direction.right) {
      currentDialPoint += distance;
    }

    while (currentDialPoint > 99) {
      currentDialPoint -= 100;
    }

    while (currentDialPoint < 0) {
      currentDialPoint += 100;
    }

    return currentDialPoint;
  }

  (Direction direction, int distance) parseInstruction(String instruction) {
    var direction = Direction.fromString(instruction.substring(0, 1));
    var distance = int.parse(instruction.substring(1));
    return (direction, distance);
  }
}

enum Direction {
  left,
  right;

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
