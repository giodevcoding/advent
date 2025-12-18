import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/ten.dart';
import 'package:test/test.dart';

void main() {
  late FactoryMachineController factoryMachineController;
  late List<String> testInput;
  setUp(() {
    factoryMachineController = FactoryMachineController();
    testInput = [
      "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
      "[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
      "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
    ];
  });

  test('FactoryMachineController optimally powers on machines', () {
    var result = factoryMachineController.turnOnMachines(testInput);

    expect(result, 7);
  });

  test('FactoryMachineController optimally configure joltage for machines', () async {
    Debugger.enable();
    var result = await factoryMachineController.configureJoltages(testInput);

    expect(result, 33);
  });

  tearDown(Debugger.disable);
}
