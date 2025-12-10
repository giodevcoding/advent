import 'dart:math';

import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';

@Day(8)
class DayEightRunner implements DayRunner {
  @override
  String run() {
    var inputReader = AdventInputReader("eight.txt");
    var boxesInput = inputReader.readIntoLines();
    var junctionBoxConnector = JunctionBoxConnector(1000);
    var circuitProduct = junctionBoxConnector.getCircuitProduct(boxesInput);
    var lastPairProduct = junctionBoxConnector.getLastPairProduct(boxesInput);
    return "\nProduct of 3 Largest Circuits: $circuitProduct"
        "\nProduct of Last Pair X Coordinates: $lastPairProduct";
  }
}

class JunctionBoxConnector {
  final int maxPairs;

  JunctionBoxConnector(this.maxPairs);

  int getLastPairProduct(List<String> input) {
    var boxes = parseInput(input);
    var pairs = getPairs(boxes);
    var (circuits, lastPair) = createCircuits(boxes, pairs);

    return lastPair.boxA.x * lastPair.boxB.x;
  }

  int getCircuitProduct(List<String> input) {
    var boxes = parseInput(input);
    var pairs = getPairs(boxes, maxPairs: maxPairs);
    var (circuits, _) = createCircuits(boxes, pairs);

    var sortedCircuits = List<Circuit>.from(
      circuits,
    )..sort((a, b) => b.junctionBoxes.length.compareTo(a.junctionBoxes.length));

    return sortedCircuits
        .sublist(0, 3)
        .fold(1, (acc, circuit) => acc * circuit.junctionBoxes.length);
  }

  (List<Circuit>, JunctionBoxPair lastPair) createCircuits(
    List<JunctionBox> boxes,
    List<JunctionBoxPair> pairs,
  ) {
    var circuits = <Circuit>[];
    late JunctionBoxPair lastPair;

    for (var pair in pairs) {
      var boxACircuitIndex = circuits.indexWhere(
        (c) => c.hasJunctionBox(pair.boxA),
      );
      var boxBCircuitIndex = circuits.indexWhere(
        (c) => c.hasJunctionBox(pair.boxB),
      );
      var boxACircuitExists = boxACircuitIndex != -1;
      var boxBCircuitExists = boxBCircuitIndex != -1;

      // 1. Neither BoxA or BoxB are in circuits => Create new circuit
      if (!boxACircuitExists && !boxBCircuitExists) {
        circuits.add(Circuit.fromPair(pair));
        lastPair = pair;
      }

      // 2. BoxA is in circuit, but not BoxB => BoxB joins circuit with BoxA
      if (boxACircuitExists && !boxBCircuitExists) {
        circuits[boxACircuitIndex].add(pair.boxB);
        lastPair = pair;
      }

      // 3. BoxB is in circuit, but not BoxA => BoxA joins circuit with BoxB
      if (!boxACircuitExists && boxBCircuitExists) {
        circuits[boxBCircuitIndex].add(pair.boxA);
        lastPair = pair;
      }

      // 4. Both BoxA and BoxB are in circuits
      if (boxACircuitExists && boxBCircuitExists) {
        // 4a They're both in the same circuit => Do Nothing
        if (boxACircuitIndex == boxBCircuitIndex) {
          continue;
        }

        // 4b They're in different circuits => Merge Circuits
        circuits[boxACircuitIndex].merge(circuits[boxBCircuitIndex]);
        circuits.removeAt(boxBCircuitIndex);
        lastPair = pair;
      }
    }

    return (circuits, lastPair);
  }

  List<JunctionBoxPair> getPairs(List<JunctionBox> boxes, {int? maxPairs}) {
    var pairs = <JunctionBoxPair>[];

    for (var i = 0; i < boxes.length - 1; i++) {
      for (var j = i + 1; j < boxes.length; j++) {
        pairs.add(JunctionBoxPair(boxes[i], boxes[j]));
      }
    }

    pairs.sort((a, b) => a.distance.compareTo(b.distance));

    if (maxPairs != null) {
      return pairs.sublist(0, maxPairs);
    }

    return pairs;
  }

  List<JunctionBox> parseInput(List<String> input) {
    return input
        .map((str) => str.split(",").map(int.parse).toList())
        .map((coords) => JunctionBox(coords[0], coords[1], coords[2]))
        .toList();
  }
}

class JunctionBox {
  final int x, y, z;
  final int id;

  static int _nextId = 0;

  JunctionBox(this.x, this.y, this.z) : id = _nextId++;

  @override
  String toString() => "JunctionBox{ id: $id, position: $x,$y,$z }";
}

class JunctionBoxPair {
  final JunctionBox boxA, boxB;
  late final num distance;

  JunctionBoxPair(this.boxA, this.boxB) {
    distance = sqrt(
      pow(boxB.x - boxA.x, 2) +
          pow(boxB.y - boxA.y, 2) +
          pow(boxB.z - boxA.z, 2),
    );
  }

  int get boxAId => boxA.id;
  int get boxBId => boxB.id;

  @override
  String toString() =>
      "JunctionBoxPair{ A: $boxA, B: $boxB, distance: $distance }";
}

class Circuit {
  final int id;
  List<JunctionBox> junctionBoxes;

  static int _nextId = 0;

  Circuit() : id = _nextId++, junctionBoxes = [];
  Circuit.fromPair(JunctionBoxPair pair)
    : id = _nextId++,
      junctionBoxes = [pair.boxA, pair.boxB];

  bool hasJunctionBox(JunctionBox box) =>
      junctionBoxes.any((b) => b.id == box.id);

  void add(JunctionBox box) => junctionBoxes.add(box);
  void merge(Circuit other) => junctionBoxes.addAll(other.junctionBoxes);

  @override
  String toString() => "Circuit { id: $id, junctionBoxes: $junctionBoxes }";
}

class JunctionBoxIdGenerator {
  JunctionBoxIdGenerator._();
}

class CircuitIdGenerator {
  static int _nextId = 0;
  static int generateId() => _nextId++;
  CircuitIdGenerator._();
}
