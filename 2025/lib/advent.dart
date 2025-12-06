import 'package:advent/src/day.dart';
import 'package:advent/src/four.dart';
import 'package:advent/src/one.dart';
import 'package:advent/src/three.dart';
import 'package:advent/src/two.dart';

String getResultForDay(int day) {
  List<DayRunner> runners = [
    DayOneRunner(),
    DayTwoRunner(),
    DayThreeRunner(),
    DayFourRunner(),
  ];
  return runners.firstWhere((r) => r.day() == day).run();
}
