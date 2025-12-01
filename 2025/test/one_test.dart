import 'package:advent/src/file_utils.dart';
import 'package:advent/src/one.dart';
import 'package:test/test.dart';

void main() {
  test(
    'NorthpoleEntrancePasswordDecoder properly decodes password from sample input with old method',
    () {
      var sampleInstructions = getSampleInput();
      var decoder = NorthPoleEntrancePasswordDecoder(startingPoint: 50);

      var result = decoder.decode(sampleInstructions);

      expect(result, 4);
    },
  );

  test(
    'NorthpoleEntrancePasswordDecoder properly decodes password from real input with old method',
    () {
      var actualInstructions = AdventInputReader("one.txt").readIntoLines();
      var decoder = NorthPoleEntrancePasswordDecoder(startingPoint: 50);

      var result = decoder.decode(actualInstructions);

      expect(result, 1034);
    },
  );

  test(
    'NorthpoleEntrancePasswordDecoder properly decodes password based touching zero',
    () {
      var sampleInstructions = getSampleInput();
      var decoder = NorthPoleEntrancePasswordDecoder(startingPoint: 50);

      var result = decoder.decode(sampleInstructions);

      expect(result, 7);
    },
  );
}

List<String> getSampleInput() {
  return ["L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82", "L32"];
}
