import 'package:advent/src/file_utils.dart';
import 'package:test/test.dart';

void main() {
  group("Test AdventInputReader.readIntoLines", () {
    test('Reads correct first line', () {
      var reader = AdventInputReader("one.txt");
      var lines = reader.readIntoLines();
      expect(lines[0], "R9");
    });

    test('Reads correct amount of lines', () {
      var reader = AdventInputReader("one.txt");
      var lines = reader.readIntoLines();
      expect(lines.length, 4186);
    });
  });
}
