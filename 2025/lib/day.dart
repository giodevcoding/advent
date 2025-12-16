import 'dart:async';

class Day {
  final int day;
  const Day(this.day);
}

abstract interface class DayRunner {
  FutureOr<String> run();
}
