import 'package:advent/src/one.dart';

String getResultForDay(int day) {
  return switch(day) {
    1 => dayOne(),
    _ => throw ArgumentError("Day does not have an existing method"),
  };
}
