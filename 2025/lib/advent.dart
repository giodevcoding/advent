import 'package:advent/day_registry.g.dart';

String getResultForDay(int day) {
  var runners = dayRunners[day] ?? [];
  return runners.map((runner) => runner.run()).join();
}
