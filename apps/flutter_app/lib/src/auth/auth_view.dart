import 'package:flutter/material.dart';

import 'auth_controller.dart';

class AuthView extends StatefulWidget {
  const AuthView({super.key, required this.controller});

  static const routeName = '/login';

  final AuthController controller;

  @override
  State<AuthView> createState() => _AuthViewState();
}

class _AuthViewState extends State<AuthView> {
  String _username = '';
  String _password = '';

  void handleLogin() {
    widget.controller.login(_username, _password);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Login'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          children: [
            TextField(
              decoration: const InputDecoration(
                hintText: 'Username',
              ),
              onChanged: (val) => _username = val,
            ),
            TextField(
              decoration: const InputDecoration(
                hintText: 'Password',
              ),
              onChanged: (val) => _password = val,
            ),
            ElevatedButton(
              onPressed: handleLogin,
              child: const Text('Login'),
            ),
          ],
        ),
      ),
    );
  }
}
