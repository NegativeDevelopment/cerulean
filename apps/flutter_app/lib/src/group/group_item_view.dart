import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';

class GroupItemView extends StatelessWidget {
  const GroupItemView({super.key, required this.groupId});

  static const routeName = "/group";

  final String groupId;

  @override
  Widget build(BuildContext context) {
    return Container(
      child: const Text('Group Item View'),
    );
  }
}
