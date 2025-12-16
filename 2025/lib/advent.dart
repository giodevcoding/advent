import 'package:advent/day_registry.g.dart';

Future<String> getResultForDay(int day) async {
  var runners = dayRunners[day] ?? [];
  var result = "";
  for (var r in runners) {
    result += await r.run();
  }

  return result;
}
