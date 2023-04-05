class GroupDebtItem {
  final String creditorId;
  final String debtorId;
  final String groupId;
  final int amount;
  final String currency;

  GroupDebtItem(
      {required this.creditorId,
      required this.debtorId,
      required this.groupId,
      required this.amount,
      required this.currency});

  factory GroupDebtItem.fromJson(Map<String, dynamic> json) {
    return GroupDebtItem(
      creditorId: json['creditorId'],
      debtorId: json['debtorId'],
      groupId: json['groupId'],
      amount: json['amount'],
      currency: json['currency'],
    );
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['creditor_id'] = creditorId;
    data['debtor_id'] = debtorId;
    data['group_id'] = groupId;
    data['amount'] = amount;
    data['currency'] = currency;
    return data;
  }
}
