import 'dart:convert';
import 'package:flutter_app/src/group/group_debt_item.dart';
import 'package:flutter_app/src/group/group_item.dart';
import 'package:flutter_app/src/group/group_transaction_item.dart';
import 'package:flutter_app/src/group/group_user_item.dart';
import 'package:http/http.dart' as http;

import '../constants.dart';

class GroupService {
  final String _token;

  GroupService(this._token);

  Map<String, String> _getFetchHeaders() {
    return <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
      'Authorization': "Bearer $_token"
    };
  }

  Future<List<GroupItem>> getGroups() async {
    var response = await http.get(Uri.parse('$appApiBaseUrl/groups'),
        headers: _getFetchHeaders());

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to load groups');
    }
  }

  Future<List<GroupItem>> getMyGroups() async {
    var response = await http.get(Uri.parse('$appApiBaseUrl/me/groups'),
        headers: _getFetchHeaders());

    if (response.statusCode == 200) {
      Iterable l = jsonDecode(response.body);
      List<GroupItem> items =
          List<GroupItem>.from(l.map((model) => GroupItem.fromJson(model)));

      return items;
    } else {
      throw Exception('Failed to load my groups');
    }
  }

  Future<List<GroupDebtItem>> getGroupDebt(String groupId) async {
    var response = await http.get(
        Uri.parse('$appApiBaseUrl/me/groups/$groupId/debts'),
        headers: _getFetchHeaders());

    if (response.statusCode == 200) {
      Iterable l = jsonDecode(response.body);
      List<GroupDebtItem> items = List<GroupDebtItem>.from(
          l.map((model) => GroupDebtItem.fromJson(model)));
      return items;
    } else {
      throw Exception('Failed to load group debt');
    }
  }

  Future<List<GroupUserItem>> getGroupMembers(String groupId) async {
    var response = await http.get(
        Uri.parse('$appApiBaseUrl/me/groups/$groupId/members'),
        headers: _getFetchHeaders());

    if (response.statusCode == 200) {
      Iterable l = jsonDecode(response.body);
      List<GroupUserItem> items = List<GroupUserItem>.from(
          l.map((model) => GroupUserItem.fromJson(model)));
      return items;
    } else {
      throw Exception('Failed to load group members');
    }
  }

  Future<List<GroupTransactionItem>> getGroupTransactions(
      String groupId) async {
    var response = await http.get(
        Uri.parse('$appApiBaseUrl/me/groups/$groupId/transactions'),
        headers: _getFetchHeaders());

    if (response.statusCode == 200) {
      Iterable l = jsonDecode(response.body);
      List<GroupTransactionItem> items = List<GroupTransactionItem>.from(
          l.map((model) => GroupTransactionItem.fromJson(model)));
      return items;
    } else {
      throw Exception('Failed to load group members');
    }
  }
}
