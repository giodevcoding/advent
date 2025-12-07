import 'dart:math';

import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';

@Day(6)
class DaySixRunner implements DayRunner {
  @override
  String run() {
    var inputReader = AdventInputReader("six.txt");
    var homeworkInput = inputReader.readIntoLines();
    var cephalopodMathSolver = CephalopodMathSolver();
    var result = cephalopodMathSolver.solve(homeworkInput);
    return "\nTotal Answer: $result";
  }
}

class CephalopodMathSolver {
  int solve(List<String> homeworkInput, {bool rtl = false}) {
    var parsedInput = rtl
        ? parseInputRTL(homeworkInput)
        : parseInput(homeworkInput);
    return parsedInput.map((mp) => mp.solve()).fold(0, (a, b) => a + b);
  }

  List<MathProblem> parseInput(List<String> homeworkInput) {
    var mathProblems = <MathProblem>[];
    var lines = homeworkInput
        .map((str) => str.split(" ").where((s) => s.trim().isNotEmpty).toList())
        .toList();

    for (var x = 0; x < lines[0].length; x++) {
      var numbers = <int>[];
      late MathOperator operator;

      for (var y = 0; y < lines.length; y++) {
        var current = lines[y][x];
        if (y == lines.length - 1) {
          operator = MathOperator.fromString(current);
          continue;
        }

        numbers.add(int.parse(current));
      }

      mathProblems.add(MathProblem(operator, numbers));
    }

    return mathProblems;
  }

  List<MathProblem> parseInputRTL(List<String> homeworkInput) {
    // TODO: Parse Input RTL
    // Use operator locations as how to identify string block start and end
    // then parse strings before any int conversion, to maintain horizontal alignment
    return [];
  }
}

class MathProblem {
  final MathOperator operator;
  final List<int> numbers;

  const MathProblem(this.operator, this.numbers);

  int solve() {
    return switch (operator) {
      MathOperator.add => numbers.reduce((a, b) => a + b),
      MathOperator.multiply => numbers.reduce((a, b) => a * b),
    };
  }
}

enum MathOperator {
  add,
  multiply;

  static MathOperator fromString(String str) {
    if (str.trim() == "*") {
      return multiply;
    }

    if (str.trim() == "+") {
      return add;
    }

    throw ArgumentError("No MathOperator for string: $str");
  }
}
