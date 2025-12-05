import 'package:advent/advent.dart' as advent;
import 'package:advent/src/debug_utils.dart';

void main(List<String> arguments) {
  Debugger.disable();
  var day = int.parse(arguments[0]);
  var result = advent.getResultForDay(day);
  print("\n----- Advent of Code - Day $day -----\n");
  print("\nResult: $result\n\n");
  print("-------------------------------------");
}
