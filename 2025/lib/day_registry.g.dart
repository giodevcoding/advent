import 'package:advent/src/five.dart';
import 'package:advent/src/four.dart';
import 'package:advent/src/one.dart';
import 'package:advent/src/three.dart';
import 'package:advent/src/two.dart';
import 'package:advent/day.dart';

final Map<int, List<DayRunner>> dayRunners = <int, List<DayRunner>>{
  1: <DayRunner>[DayOneRunner()],
  2: <DayRunner>[DayTwoRunner()],
  3: <DayRunner>[DayThreeRunner()],
  4: <DayRunner>[DayFourRunner()],
  5: <DayRunner>[DayFiveRunner()],
};
