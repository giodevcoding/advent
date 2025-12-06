import 'package:advent/day.dart';
import 'package:advent/src/debug_utils.dart';
import 'package:advent/src/file_utils.dart';
import 'package:advent/src/math_utils.dart';
import 'package:advent/src/string_utils.dart';

@Day(2)
class DayTwoRunner implements DayRunner {
  @override
  String run() {
    var fileReader = AdventInputReader("two.txt");
    var productIds = fileReader.read();
    var sValidator = SingleSplitValidator();
    var singleSplit = sValidator.sumInvalidProductIds(productIds).toString();
    var mValidator = MultiSplitValidator();
    var multiSplit = mValidator.sumInvalidProductIds(productIds).toString();
    return (singleSplit: singleSplit, multiSplit: multiSplit).toString();
  }
}

abstract class ProductIdValidator {
  int getInvalidIdValue(int id);

  int sumInvalidProductIds(String productIds) {
    var invalidProductIdsSum = 0;
    var ids = parseProductIds(productIds);
    for (var id in ids) {
      invalidProductIdsSum += getInvalidIdValue(id);
    }
    return invalidProductIdsSum;
  }

  List<int> parseProductIds(String productIds) {
    return productIds
        .split(",")
        .map((r) => r.split("-"))
        .map((r) => r.map(int.parse).toList())
        .expand((r) => getIdsFromRange(r[0], r[1]))
        .toList();
  }

  List<int> getIdsFromRange(int low, int high) {
    List<int> ids = [];
    for (var id = low; id <= high; id++) {
      ids.add(id);
    }
    return ids;
  }
}

class SingleSplitValidator extends ProductIdValidator {
  @override
  int getInvalidIdValue(int id) {
    var idString = id.toString();
    var idLength = idString.length;

    if (idLength % 2 != 0) {
      return 0;
    }

    var left = idString.substring(0, idLength ~/ 2);
    var right = idString.substring(idLength ~/ 2);

    if (left != right) {
      return 0;
    }

    return id;
  }
}

class MultiSplitValidator extends ProductIdValidator {
  @override
  int getInvalidIdValue(int id) {
    var idString = id.toString();
    var divs = divisors(idString.length).where((div) => div != idString.length);
    Debugger.log(divs);

    for (var div in divs) {
      var chunks = chunk(idString, div);
      Debugger.log(chunks);
      Debugger.waitForInput();
      if (areAllPartsEqual(chunks)) {
        Debugger.log("Invalid: $id");
        Debugger.waitForInput();
        return id;
      }
    }

    return 0;
  }

  bool areAllPartsEqual(List<dynamic> parts) {
    for (var i = 0; i < parts.length - 1; i++) {
      if (parts[i] != parts[i + 1]) {
        return false;
      }
    }

    return true;
  }
}
