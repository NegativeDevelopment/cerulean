import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:http/http.dart' as http;

import '../common/model/config.dart';
import '../common/model/token.dart';

// Create a stateful widget for the login page
class LoginPage extends StatefulWidget {
  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  // Create controllers for email and password fields
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        alignment: Alignment.center,
        padding: const EdgeInsets.all(32.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              'Welcome',
              style: Theme.of(context).textTheme.headline2,
            ),
            const SizedBox(height: 48),
            SizedBox(
              width: 300,
              child: TextField(
                controller: usernameController,
                decoration: const InputDecoration(
                  labelText: 'Username',
                  border: OutlineInputBorder(
                      borderRadius: BorderRadius.all(Radius.circular(4))),
                ),
              ),
            ),
            const SizedBox(height: 8),
            SizedBox(
              width: 300,
              child: TextField(
                controller: passwordController,
                obscureText: true,
                decoration: const InputDecoration(
                  labelText: 'Password',
                  border: OutlineInputBorder(
                      borderRadius: BorderRadius.all(Radius.circular(4))),
                ),
              ),
            ),
            const SizedBox(height: 24),
            ElevatedButton(
              onPressed: (() {
                String username = usernameController.text;
                String password = passwordController.text;

                if (username.isEmpty || password.isEmpty) return;

                login(username, password);
              }),
              child: const Text('Login'),
            ),
            TextButton(
              onPressed: () {
                String username = usernameController.text;
                String password = passwordController.text;

                if (username.isEmpty || password.isEmpty) return;

                register(username, password);
              },
              child: const Text('Register'),
            ),
          ],
        ),
      ),
    );
  }

  login(String username, String password) async {
    var tokenModel = Provider.of<TokenModel>(context, listen: false);
    var appConfig = Provider.of<ConfigModel>(context, listen: false);

    var loginUri = Uri.parse('${appConfig.apiUri}/auth/login');

    // Make the HTTP POST request to get the token
    var response = await http.post(
      loginUri,
      body: json.encode({'username': username, 'password': password}),
    );

    // If the request was successful, parse the JSON and get the token
    if (response.statusCode == 200) {
      var json = jsonDecode(response.body);
      String token = json['token'];

      tokenModel.setToken(token);
    }
  }

  register(String username, String password) async {
    if (!mounted) return;
    var appConfig = Provider.of<ConfigModel>(context, listen: false);

    var registerUri = Uri.parse('${appConfig.apiUri}/auth/register');

    // Make the HTTP POST request to get the token
    var response = await http.post(
      registerUri,
      body: json.encode({'username': username, 'password': password}),
    );

    // If the request was successful, parse the JSON and get the token
    if (response.statusCode == 200) {
      var json = jsonDecode(response.body);
      debugPrint(json);
    }
  }

  clearForm() {
    usernameController.clear();
    passwordController.clear();
  }
}
