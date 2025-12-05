import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/two.dart';
import 'package:test/test.dart';

void main() {
  late String testInput;
  late ProductIdValidator validator;

  setUp(() {
    testInput =
        "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124";
  });

  group('SplitSplitValidator', () {
    setUp(() {
      validator = SingleSplitValidator();
    });

    test('sumInvalidProductIds gets correct sum on invalid product ids', () {
      var result = validator.sumInvalidProductIds(testInput);
      expect(result, 1227775554);
    });
  });

  group('MultiSplitValidator', () {
    setUp(() {
      validator = MultiSplitValidator();
    });

    test('sumInvalidProductIds gets correct sum on invalid product ids', () {
      var result = validator.sumInvalidProductIds(testInput);
      expect(result, 4174379265);
    });
  });

  tearDown(() {
    Debugger.disable();
  });
}
