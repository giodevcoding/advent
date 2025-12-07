import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/six.dart';
import 'package:test/test.dart';

void main() {
  late List<String> testInput;
  late CephalopodMathSolver cephalopodMathSolver;

  setUp(() {
    testInput = [
      "123 328  51 64 ",
      " 45 64  387 23 ",
      "  6 98  215 314",
      "*   +   *   +",
    ];
    cephalopodMathSolver = CephalopodMathSolver();
  });

  test('CephalopodMathSolver.solve solves math correctly', () {
    var result = cephalopodMathSolver.solve(testInput);
    expect(result, 4277556);
  });

  test('CephalopodMathSolver.solve with RTL solves math correctly', () {
    Debugger.enable();
    var result = cephalopodMathSolver.solve(testInput, rtl: true);
    expect(result, 3263827);
  });

  tearDown(Debugger.disable);
}
