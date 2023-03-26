import 'package:flutter/widgets.dart';
import 'package:flutter_app/src/auth/auth_controller.dart';
import 'package:flutter_app/src/group/group_item.dart';
import 'package:flutter_app/src/group/group_service.dart';

import '../auth/auth_service.dart';

class GroupController with ChangeNotifier {
  GroupController(AuthController authController)
      : _groupService = GroupService(authController.token) {
    loadGroups();
  }

  final GroupService _groupService;

  List<GroupItem> userGroups = [];

  Future<void> loadGroups() async {
    userGroups = await _groupService.getMyGroups();
    notifyListeners();
  }
}
