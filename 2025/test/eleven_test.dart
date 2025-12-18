import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/eleven.dart';
import 'package:test/test.dart';

void main() {
  late ReactorDeviceManager reactorDeviceManager;
  late List<String> testInput;
  late List<String> testInput2;

  setUp(() {
    testInput = [
      "aaa: you hhh",
      "you: bbb ccc",
      "bbb: ddd eee",
      "ccc: ddd eee fff",
      "ddd: ggg",
      "eee: out",
      "fff: out",
      "ggg: out",
      "hhh: ccc fff iii",
      "iii: out",
    ];

    testInput2 = [
      "svr: aaa bbb",
      "aaa: fft",
      "fft: ccc",
      "bbb: tty",
      "tty: ccc",
      "ccc: ddd eee",
      "ddd: hub",
      "hub: fff",
      "eee: dac",
      "dac: fff",
      "fff: ggg hhh",
      "ggg: out",
      "hhh: out",
    ];
  });

  test('ReactorDeviceManager finds number of routes from you to out', () {
    // Debugger.enable();
    reactorDeviceManager = ReactorDeviceManager(testInput);
    var result = reactorDeviceManager.countAllPaths("you");

    expect(result, 5);
  });

  test('ReactorDeviceManager finds number of routes that pass dac and fft', () {
    Debugger.enable();
    reactorDeviceManager = ReactorDeviceManager(
      testInput2,
      cache: DevicePassReactorCache(),
    );
    var result = reactorDeviceManager.countAllPaths("svr");

    expect(result, 2);
  });

  tearDown(Debugger.disable);
}
