import 'package:flutter/material.dart';

class GroupCardWidget extends StatelessWidget {
  final String groupName;
  final int memberCount;

  const GroupCardWidget(
      {Key? key, required this.groupName, required this.memberCount})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 12.0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  groupName,
                  style: Theme.of(context).textTheme.headline6,
                ),
                Text(
                  "$memberCount members",
                  style: Theme.of(context).textTheme.bodyText2,
                ),
                // Text(
                //   "50 Transactions",
                //   style: Theme.of(context).textTheme.bodyText2,
                // ),
              ],
            ),
            ClipRRect(
              borderRadius: const BorderRadius.all(Radius.circular(8.0)),
              child: Image.asset(
                "assets/group-placeholder.jpg",
                width: 60,
                height: 60,
                fit: BoxFit.cover,
              ),
            )
          ],
        ),
      ),
    );
  }
}
