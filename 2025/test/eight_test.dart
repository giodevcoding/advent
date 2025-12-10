import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/eight.dart';
import 'package:test/test.dart';

void main() {
  late JunctionBoxConnector junctionBoxConnector;
  late List<String> testInput;

  setUp(() {
    junctionBoxConnector = JunctionBoxConnector(10);
    testInput = [
      "162,817,812",
      "57,618,57",
      "906,360,560",
      "592,479,940",
      "352,342,300",
      "466,668,158",
      "542,29,236",
      "431,825,988",
      "739,650,466",
      "52,470,668",
      "216,146,977",
      "819,987,18",
      "117,168,530",
      "805,96,715",
      "346,949,466",
      "970,615,88",
      "941,993,340",
      "862,61,35",
      "984,92,344",
      "425,690,689",
    ];
  });

  test('JunctionBoxConnector returns correct circuits', () {
    Debugger.enable();

    var result = junctionBoxConnector.getCircuitProduct(testInput);

    expect(result, 40);
  });

  test('JunctionBoxConnector returns correct last pair', () {
    Debugger.enable();

    var result = junctionBoxConnector.getLastPairProduct(testInput);

    expect(result, 25272);
  });

  tearDown(Debugger.disable);
}
