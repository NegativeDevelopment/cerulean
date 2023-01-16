import 'package:flutter/material.dart';

class ConfigModel with ChangeNotifier {
  String get apiUri {
    bool isDev = !const bool.fromEnvironment('dart.vm.product');
    if (isDev) {
      return "http://localhost:3000/api";
    }

    return "https://example.com";
  }
}
