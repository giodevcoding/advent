import 'package:advent/src/day.dart';
import 'package:advent/src/one.dart';
import 'package:advent/src/two.dart';

String getResultForDay(int day) {
  List<DayRunner> runners = [DayOneRunner(), DayTwoRunner()];
  return runners.firstWhere((r) => r.day() == day).run();
}
