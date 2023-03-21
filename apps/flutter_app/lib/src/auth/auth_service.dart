import 'dart:convert';

import 'package:flutter_app/src/constants.dart';
import 'package:http/http.dart' as http;

class AuthService {
  Future<String> login(String username, String password) async {
    var response = await http.post(
      Uri.parse('$appApiBaseUrl/auth/login'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, String>{
        'username': username,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body)['token'];
    } else {
      throw Exception('Failed to Login');
    }
  }

  Future<bool> register(String username, String password) async {
    return false;
  }
}
