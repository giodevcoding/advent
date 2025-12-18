import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';
import 'package:advent/src/math_utils.dart';

@Day(11)
class DayElevenRunner implements DayRunner {
  @override
  String run() {
    Debugger.enable();
    var input = AdventInputReader("eleven.txt").readIntoLines();
    var reactorDeviceManager = ReactorDeviceManager(input);
    var pathCountYou = reactorDeviceManager.countAllPaths("you");
    var reactorDeviceManagerWithPassCache = ReactorDeviceManager(
      input,
      cache: DevicePassReactorCache(),
    );
    var pathCountSvr = reactorDeviceManagerWithPassCache.countAllPaths("svr");
    return "\nPath Count From You to Out: $pathCountYou"
        "\nPath Count with Correct Passes: $pathCountSvr";
  }
}

class ReactorDeviceManager {
  final ReactorCache cache;
  late final Map<String, List<String>> deviceGraph;

  ReactorDeviceManager(List<String> input, {ReactorCache? cache})
    : cache = cache ?? StatelessReactorCache() {
    deviceGraph = getDeviceGraphMap(input);
  }

  Map<String, List<String>> getDeviceGraphMap(List<String> input) {
    var deviceGraphMap = <String, List<String>>{};
    for (var line in input) {
      var split = line.split(":");
      deviceGraphMap[split[0]] = split[1]
          .split(" ")
          .where((str) => str.trim().isNotEmpty)
          .toList();
    }
    return deviceGraphMap;
  }

  int countAllPaths(String node, [ReactorCacheState? state]) {
    var newState = cache.generateState(node, state);

    if (cache.exists(node, newState)) {
      return cache.get(node, newState);
    }

    if (node == "out") {
      if (cache.isStateSuccessful(newState)) {
        return 1;
      } else {
        return 0;
      }
    }

    var result = deviceGraph[node]!
        .map((child) => countAllPaths(child, newState))
        .reduce(sum);

    cache.save(node, newState, result);

    return result;
  }
}

abstract class ReactorCache<T> {
  bool exists(String node, ReactorCacheState<T> state) =>
      _internalCache.containsKey((node, state));
  int save(String node, ReactorCacheState<T> state, int value) =>
      _internalCache.putIfAbsent((node, state), () => value);
  int get(String node, ReactorCacheState<T> state) =>
      _internalCache[(node, state)]!;

  Map<(String node, ReactorCacheState<T> state), int> get _internalCache;
  ReactorCacheState<T> generateState(
    String node,
    ReactorCacheState<T>? oldState,
  );
  bool isStateSuccessful(ReactorCacheState<T> state);
}

class StatelessReactorCache extends ReactorCache<bool> {
  final _cache = <(String, ReactorCacheState<bool>), int>{};

  @override
  Map<(String, ReactorCacheState<bool>), int> get _internalCache => _cache;

  @override
  ReactorCacheState<bool> generateState(String node, _) =>
      ReactorCacheState(true);

  @override
  int save(String node, ReactorCacheState<bool> state, int value) =>
      _cache[(node, state)] = value;

  @override
  bool isStateSuccessful(ReactorCacheState<bool> state) => true;
}

class DevicePassReactorCache
    extends ReactorCache<({bool passedDac, bool passedFft})> {
  final _cache =
      <(String, ReactorCacheState<({bool passedDac, bool passedFft})>), int>{};

  @override
  Map<(String, ReactorCacheState<({bool passedDac, bool passedFft})>), int>
  get _internalCache => _cache;

  @override
  ReactorCacheState<({bool passedDac, bool passedFft})> generateState(
    String node,
    ReactorCacheState<({bool passedDac, bool passedFft})>? oldState,
  ) {
    var passedDac = node == "dac" ? true : oldState?.value.passedDac ?? false;
    var passedFft = node == "fft" ? true : oldState?.value.passedFft ?? false;
    return ReactorCacheState((passedDac: passedDac, passedFft: passedFft));
  }

  @override
  bool isStateSuccessful(
    ReactorCacheState<({bool passedDac, bool passedFft})> state,
  ) => state.value.passedDac && state.value.passedFft;
}

class ReactorCacheState<T> {
  final T value;
  const ReactorCacheState(this.value);

  ReactorCacheState<T> clone() => ReactorCacheState(value);

  @override
  bool operator ==(Object other) {
    if (other is T) {
      return other == value;
    }

    return other is ReactorCacheState<T> &&
        other.runtimeType == runtimeType &&
        other.value == value;
  }

  @override
  int get hashCode => value.hashCode;
}
