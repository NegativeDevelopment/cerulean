import 'package:flutter/material.dart';
import 'package:flutter_app/src/auth/auth_controller.dart';
import 'package:flutter_app/src/group/group_item_controller.dart';

class GroupItemView extends StatelessWidget {
  GroupItemView(
      {super.key, required this.authController, required this.groupId})
      : _groupItemController = GroupItemController(authController, groupId);

  static const routeName = "/group";

  final GroupItemController _groupItemController;
  final AuthController authController;
  final String groupId;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Individual Group'),
      ),
      body: AnimatedBuilder(
        // [AnimatedBuilder] accepts any [Listenable] subtype.
        animation: _groupItemController,
        builder: (BuildContext context, Widget? child) {
          return Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              SizedBox(height: 16.0),
              const Text(
                'Debts',
                style: TextStyle(
                  fontSize: 20.0,
                  fontWeight: FontWeight.bold,
                ),
              ),
              Expanded(
                child: ListView.builder(
                  itemCount: _groupItemController.debts.length,
                  itemBuilder: (BuildContext context, int index) {
                    return ListTile(
                      title: Text(_groupItemController.debts[index].currency),
                      subtitle: Text(
                          _groupItemController.debts[index].amount.toString()),
                    );
                  },
                ),
              ),
              const SizedBox(height: 16.0),
              const Text(
                'Transactions',
                style: TextStyle(
                  fontSize: 20.0,
                  fontWeight: FontWeight.bold,
                ),
              ),
              Expanded(
                child: ListView.builder(
                  itemCount: _groupItemController.transactions.length,
                  itemBuilder: (BuildContext context, int index) {
                    return ListTile(
                      title: Text(
                          _groupItemController.transactions[index].description),
                      subtitle: Text(_groupItemController
                          .transactions[index].amount
                          .toString()),
                    );
                  },
                ),
              ),
            ],
          );
        },
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          // Add new entry logic
        },
        child: Icon(Icons.add),
      ),
    );
  }
}
