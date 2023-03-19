import 'package:flutter/material.dart';

import 'auth_service.dart';

class AuthController with ChangeNotifier {
  AuthController(this._authService);

  final AuthService _authService;

  bool _isLoggedIn = false;

  void login(String username, String password) async {
    _isLoggedIn = await _authService.login(username, password);
    notifyListeners();
  }
}
