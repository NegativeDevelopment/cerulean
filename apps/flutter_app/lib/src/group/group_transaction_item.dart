import 'package:flutter/material.dart';
import 'package:flutter_app/src/group/group_transaction_member_item.dart';
import 'package:flutter_app/src/group/group_user_item.dart';
import 'package:uuid/uuid.dart';

class GroupTransactionItem {
  final String id;
  final DateTime createdAt;
  final DateTime updatedAt;
  final DateTime? deletedAt;
  final String groupId;
  final double amount;
  final String currency;
  final String title;
  final String description;
  final String createdByUserId;
  final GroupUserItem createdBy;
  final List<GroupTransactionMemberItem> members;

  GroupTransactionItem({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    this.deletedAt,
    required this.groupId,
    required this.amount,
    required this.currency,
    required this.title,
    required this.description,
    required this.createdByUserId,
    required this.createdBy,
    required this.members,
  });

  factory GroupTransactionItem.fromJson(Map<String, dynamic> json) {
    final List<dynamic> membersJson = json['members'];
    final List<GroupTransactionMemberItem> members =
        membersJson.map((e) => GroupTransactionMemberItem.fromJson(e)).toList();
    return GroupTransactionItem(
      id: json['id'],
      groupId: json['group_id'],
      amount: json['amount'],
      currency: json['currency'],
      title: json['title'],
      description: json['description'],
      createdByUserId: json['created_by_user_id'],
      createdBy: GroupUserItem.fromJson(json['created_by']),
      members: members,
      createdAt: DateTime.parse(json['created_at']),
      updatedAt: DateTime.parse(json['updated_at']),
      deletedAt: json['deleted_at'] != null
          ? DateTime.parse(json['deleted_at'])
          : null,
    );
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['id'] = id;
    data['created_at'] = createdAt.toIso8601String();
    data['updated_at'] = updatedAt.toIso8601String();
    if (deletedAt != null) data['deleted_at'] = deletedAt!.toIso8601String();
    data['group_id'] = groupId;
    data['amount'] = amount;
    data['currency'] = currency;
    data['title'] = title;
    data['description'] = description;
    data['created_by_user_id'] = createdByUserId;
    data['created_by'] = createdBy.toJson();
    data['members'] = members.map((e) => e.toJson()).toList();
    return data;
  }
}
