import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';

class GroupListView extends StatelessWidget {
  const GroupListView({super.key});

  static const routeName = "/";

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Group List'),
      ),
      body: const Center(
        child: Text('Group List'),
      ),
    );
  }
}
