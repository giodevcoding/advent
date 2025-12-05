import 'dart:io';

class Debugger {
  static bool _enabled = false;

  Debugger._();

  static void log (Object? object) {
    if (_enabled) {
      print(object);
    }
  }

  static void waitForInput({ bool? when }) {
    if (_enabled && (when ?? true)) {
      stdin.readLineSync();
    }
  }

  static void enable() {
    _enabled = true;
  }

  static void disable() {
    _enabled = false;
  }

  static get enabled => _enabled;
}
