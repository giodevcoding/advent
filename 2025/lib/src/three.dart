import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';

@Day(3)
class DayThreeRunner implements DayRunner {
  @override
  String run() {
    var fileReader = AdventInputReader("three.txt");
    var banksInput = fileReader.readIntoLines();
    var joltageReader = JoltageReader();
    var twoBatteryTotalVoltageOutput = joltageReader.getTotalJoltageOutput(
      banksInput,
      2,
    );
    var twelveBatteryTotalVoltageOutput = joltageReader.getTotalJoltageOutput(
      banksInput,
      12,
    );
    return "\n2 Batteries: $twoBatteryTotalVoltageOutput\n12 Batteries: $twelveBatteryTotalVoltageOutput";
  }
}

class JoltageReader {
  int getTotalJoltageOutput(List<String> banksInput, int maxBatteries) {
    var banks = _parseBanks(banksInput);
    var total = 0;
    for (var bank in banks) {
      var voltage = _calculateMaxBankVoltage(bank, maxBatteries);
      Debugger.log("bank: $bank = voltage: $voltage");
      Debugger.waitForInput();
      total += voltage;
    }

    return total;
  }

  List<List<int>> _parseBanks(List<String> banksInput) {
    return banksInput
        .map((bank) => bank.split("").map(int.parse).toList())
        .toList();
  }

  int _calculateMaxBankVoltage(List<int> bank, int maxBatteries) {
    var digits = <int>[];
    var latestIndex = -1;
    for (var i = 0; i < maxBatteries; i++) {
      var batteriesLeft = maxBatteries - i;
      var end = bank.length - batteriesLeft;
      var foundIndex = _findNextLargestDigitIndexInRange(
        bank,
        latestIndex + 1,
        end,
      );
      Debugger.log((
        start: latestIndex + 1,
        end: end,
        found: bank[foundIndex],
        batteriesLeft: batteriesLeft,
      ));
      Debugger.waitForInput();
      digits.add(bank[foundIndex]);
      latestIndex = foundIndex;
    }

    return int.parse(digits.map((d) => d.toString()).join());
  }

  int _findNextLargestDigitIndexInRange(List<int> digits, int start, int end) {
    var largestIndex = start;
    for (var i = start; i <= end; i++) {
      var current = digits[i];
      if (current == 9) {
        return i;
      }

      if (current > digits[largestIndex]) {
        largestIndex = i;
      }
    }
    return largestIndex;
  }
}
