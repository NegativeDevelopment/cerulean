class GroupItem {
  final String id;
  final String name;
  final String ownerID;
  // final List<Member> members;
  // final List<Transaction> transactions;
  // final List<Debt> debts;

  GroupItem({
    required this.id,
    required this.name,
    required this.ownerID,
    // this.members,
    // this.transactions,
    // this.debts,
  });

  factory GroupItem.fromJson(Map<String, dynamic> json) {
    return GroupItem(
      id: json['id'],
      name: json['name'],
      ownerID: json['owner_id'],
      // members: json['members'],
      // transactions: json['transactions'],
      // debts: json['debts'],
    );
  }
}
