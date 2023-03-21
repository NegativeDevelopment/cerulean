import 'package:flutter/widgets.dart';

import '../auth/auth_service.dart';

class GroupController with ChangeNotifier {
  GroupController(this._authService);

  final AuthService _authService;
}
