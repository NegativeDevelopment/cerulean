import 'package:flutter/material.dart';

class TokenModel with ChangeNotifier {
  String? _token;

  String? get token => _token;

  setToken(String token) {
    _token = token;
    notifyListeners();
  }

  isLoggedIn() => _token != null;
}
