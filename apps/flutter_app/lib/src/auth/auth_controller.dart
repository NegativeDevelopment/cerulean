import 'package:flutter/material.dart';

import 'auth_service.dart';

class AuthController with ChangeNotifier {
  AuthController(this._authService);

  final AuthService _authService;

  late String token;
  bool _isLoggedIn = false;

  void login(String username, String password) async {
    var token = await _authService.login(username, password);
    if (token.isEmpty) {
      _isLoggedIn = false;
    } else {
      _isLoggedIn = true;
    }

    this.token = token;
    notifyListeners();
  }
}
