import 'dart:convert';
import 'dart:io';
import 'dart:math';

import 'package:advent/day.dart';
import 'package:advent/src/file_utils.dart';
import 'package:advent/src/math_utils.dart';
import 'package:advent/src/string_utils.dart';

import 'debug_utils.dart';

@Day(10)
class DayTenRunner implements DayRunner {
  @override
  Future<String> run() async {
    var inputReader = AdventInputReader("ten.txt");
    var input = inputReader.readIntoLines();
    var factoryMachineController = FactoryMachineController();
    var minPressesTurnOnSum = factoryMachineController.turnOnMachines(input);
    var minPressesConfigureJoltageSum = await factoryMachineController
        .configureJoltages(input);
    return "\nMinimum Presses to Turn On Sum: $minPressesTurnOnSum"
        "\nMinimum Presses to Configure Joltage: $minPressesConfigureJoltageSum";
  }
}

class FactoryMachineController {
  int turnOnMachines(List<String> machineInput) {
    return machineInput
        .map(MachineDescription.fromString)
        .map(getMinimumButtonPresses)
        .reduce(sum);
  }

  Future<int> configureJoltages(List<String> machineInput) async {
    var machines = machineInput.map(MachineDescription.fromString).toList();
    var totalPresses = 0;
    for (var machine in machines) {
      totalPresses += await configureJoltage(machine);
    }

    return totalPresses;
  }

  Future<int> configureJoltage(MachineDescription machine) async {
    try {
      final process = await Process.start("glpsol", [
        '--lp',
        '/dev/stdin',
        '-o',
        '/dev/stdout',
      ]);

      process.stdin.write(machine.toLPFormat());
      await process.stdin.close();

      final stdout = await process.stdout.transform(utf8.decoder).join();

      // Wait for process to complete
      final exitCode = await process.exitCode;

      if (exitCode != 0) {
        final stderr = await process.stderr.transform(utf8.decoder).join();
        throw Exception('glpsol failed: $stderr');
      }

      var solutionStr = RegExp(r"Objective:\s+obj\s+=\s+(\d+)").firstMatch(stdout)?.group(1);
      return int.parse(solutionStr!);
    } catch (err) {
      throw Exception('glpsol failed: $err');
    }
  }

  int getMinimumButtonPresses(MachineDescription machine) {
    return getAllValidButtonPressCombos(
      machine,
    ).fold(maxInt, (smallest, current) => min(smallest, current.reduce(sum)));
  }

  List<List<int>> getAllValidButtonPressCombos(
    MachineDescription machine, {
    List<List<int>> validPresses = const [],
    List<int> currentPresses = const [],
    index = 0,
  }) {
    var presses = index == 0 ? <int>[] : currentPresses;
    var vPresses = index == 0 ? <List<int>>[] : validPresses;

    if (presses.length < machine.buttonWirings.length) {
      getAllValidButtonPressCombos(
        machine,
        validPresses: vPresses,
        currentPresses: [...presses, 1],
        index: index + 1,
      );
      getAllValidButtonPressCombos(
        machine,
        validPresses: vPresses,
        currentPresses: [...presses, 0],
        index: index + 1,
      );
    } else {
      var result = List<int>.filled(machine.indicatorLights.length, 0);
      for (var i = 0; i < presses.length; i++) {
        if (presses[i] == 1) {
          var button = machine.buttonWirings[i];
          for (var j = 0; j < result.length; j++) {
            if (button.contains(j)) {
              result[j] = (result[j] + 1) % 2;
            }
          }
        }
      }
      var isValid = true;
      for (var i = 0; i < result.length; i++) {
        if (result[i] != machine.indicatorLights[i]) {
          isValid = false;
          break;
        }
      }
      if (isValid) {
        vPresses.add(presses);
      }
    }

    return vPresses;
  }
}

class MachineDescription {
  late final List<int> indicatorLights;
  late final List<List<int>> buttonWirings;
  late final List<int> joltageRequirements;

  MachineDescription.fromString(String input) {
    indicatorLights = input
        .substring(1, input.indexOf("]"))
        .split('')
        .map((c) => c == '#' ? 1 : 0)
        .toList();

    buttonWirings = input
        .substring(input.indexOf("("), input.lastIndexOf(")") + 1)
        .split(" ")
        .map(
          (wiring) => wiring
              .substring(1, wiring.length - 1)
              .split(",")
              .map(int.parse)
              .toList(),
        )
        .toList();

    joltageRequirements = input
        .substring(input.indexOf("{") + 1, input.length - 1)
        .split(",")
        .map(int.parse)
        .toList();
  }

  @override
  String toString() {
    var indicatorLightsString =
        "[${indicatorLights.map((l) => l == 1 ? '#' : '.').join()}]";
    var buttonWiringsString = buttonWirings
        .map((wiring) => "(${wiring.join(',')})")
        .join(" ");
    var joltageRequirementsString = "{${joltageRequirements.join(',')}}";
    return "$indicatorLightsString $buttonWiringsString $joltageRequirementsString";
  }

  String toLPFormat() {
    var minimizeInfo = buttonWirings
        .asMap()
        .keys
        .map((k) => alphabet[k])
        .join(" + ");
    var boundsInfo = buttonWirings
        .asMap()
        .keys
        .map((k) => "${alphabet[k]} >= 0")
        .toList();
    var generalInfo = buttonWirings
        .asMap()
        .keys
        .map((k) => alphabet[k])
        .join(" ");
    var subjectToInfo = <String>[];
    for (var i = 0; i < joltageRequirements.length; i++) {
      var relevantButtonIndices = buttonWirings
          .where((b) => b.contains(i))
          .map(buttonWirings.indexOf)
          .map((i) => alphabet[i])
          .toList();

      subjectToInfo.add(
        "${relevantButtonIndices.join(' + ')} = ${joltageRequirements[i]}",
      );
    }

    return """
Minimize
  obj: $minimizeInfo

Subject To
${subjectToInfo.map((s) => "  $s").join("\n")}

Bounds
${boundsInfo.map((s) => "  $s").join("\n")}

General
  $generalInfo

End
""";
  }
}
