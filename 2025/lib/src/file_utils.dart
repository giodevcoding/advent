import 'dart:io';

class AdventInputReader {
  final String filename;
  static const filePath = "lib/src/inputs/";

  AdventInputReader(this.filename);

  List<String> readIntoLines() {
    return File(filePath + filename).readAsLinesSync();
  }

  String read() {
    return File(filePath + filename).readAsStringSync();
  }
}
