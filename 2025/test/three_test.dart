import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/three.dart';
import 'package:test/test.dart';

void main() {
  late JoltageReader joltageReader;
  var testInput = [
    "987654321111111",
    "811111111111119",
    "234234234234278",
    "818181911112111",
  ];

  setUp(() {
    joltageReader = JoltageReader();
  });

  test(
    'JoltageReader.getTotalJoltageOutput with test input and 2 batteries produces expected output',
    () {
      var batteryCount = 2;
      var totalVoltageOutput = joltageReader.getTotalJoltageOutput(
        testInput,
        batteryCount,
      );
      expect(totalVoltageOutput, 357);
    },
  );

  test(
    'JoltageReader.getTotalJoltageOutput with test input and 12 batteries produces expected output',
    () {
      var batteryCount = 12;
      var totalVoltageOutput = joltageReader.getTotalJoltageOutput(
        testInput,
        batteryCount,
      );
      expect(totalVoltageOutput, 3121910778619);
    },
  );

  tearDown(() {
    Debugger.disable();
  });
}
