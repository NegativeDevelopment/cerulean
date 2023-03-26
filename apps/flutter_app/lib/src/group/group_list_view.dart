import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:flutter_app/src/auth/auth_controller.dart';
import 'package:flutter_app/src/group/group_controller.dart';

class GroupListView extends StatelessWidget {
  GroupListView({super.key, required this.authController})
      : _groupController = GroupController(authController);

  static const routeName = "/";

  final AuthController authController;
  final GroupController _groupController;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Group List'),
      ),
      body: AnimatedBuilder(
        // [AnimatedBuilder] accepts any [Listenable] subtype.
        animation: _groupController,
        builder: (BuildContext context, Widget? child) {
          return ListView.builder(
            // Let the ListView know how many items it needs to build.
            itemCount: _groupController.userGroups.length,
            // Provide a builder function. This is where the magic happens.
            // Convert each item into a widget based on the type of item it is.
            itemBuilder: (context, index) {
              final item = _groupController.userGroups[index];

              return ListTile(
                title: Text(item.name),
                subtitle: Text(item.id),
              );
            },
          );
        },
      ),
    );
  }
}
