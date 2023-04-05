import 'package:flutter/widgets.dart';
import 'package:flutter_app/src/group/group_debt_item.dart';
import 'package:flutter_app/src/group/group_service.dart';
import 'package:flutter_app/src/group/group_transaction_item.dart';

import '../auth/auth_controller.dart';

class GroupItemController with ChangeNotifier {
  GroupItemController(AuthController authController, this._groupId)
      : _groupService = GroupService(authController.token) {
    loadGroup();
  }

  final GroupService _groupService;
  final String _groupId;

  List<GroupTransactionItem> transactions = [];
  List<GroupDebtItem> debts = [];

  Future<void> loadGroup() async {
    transactions = await _groupService.getGroupTransactions(_groupId);
    debts = await _groupService.getGroupDebt(_groupId);
    notifyListeners();
  }
}
