import 'package:cerulean_app/common/widgets/group_card_widget.dart';
import 'package:flutter/material.dart';

import '../common/model/group.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          padding: const EdgeInsets.all(16.0),
          child: Text(
            "Group List",
            style: Theme.of(context).textTheme.headline5,
          ),
        ),
        Expanded(
          child: ListView(
              children: _getGroups()
                  .map((group) => GroupCardWidget(
                      groupName: group.name, memberCount: group.memberCount))
                  .toList()),
        ),
      ],
    );
  }

  // Retrive list of groups from backend api
  List<Group> _getGroups() {
    return [
      Group(
        name: "Group 1",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 2",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 3",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 4",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 5",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 6",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 7",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 8",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 9",
        memberCount: 102,
        transactionCount: 50,
      ),
      Group(
        name: "Group 10",
        memberCount: 102,
        transactionCount: 50,
      ),
    ];
  }
}
