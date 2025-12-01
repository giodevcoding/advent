import 'package:advent/advent.dart' as advent;

void main(List<String> arguments) {
  var day = int.parse(arguments[0]);
  var result = advent.getResultForDay(day);
  print("\n----- Advent of Code - Day $day -----\n");
  print("\nResult: $result\n\n");
  print("-------------------------------------");
}
