import 'package:flutter/foundation.dart';
import 'package:flutter_app/src/group/group_user_item.dart';
import 'package:uuid/uuid.dart';

class GroupTransactionMemberItem {
  final String transactionId;
  final GroupUserItem user;
  final DateTime createdAt;
  final DateTime updatedAt;
  final DateTime? deletedAt;

  GroupTransactionMemberItem({
    required this.transactionId,
    required this.user,
    required this.createdAt,
    required this.updatedAt,
    this.deletedAt,
  });

  factory GroupTransactionMemberItem.fromJson(Map<String, dynamic> json) {
    return GroupTransactionMemberItem(
      transactionId: json['transaction_id'],
      user: GroupUserItem.fromJson(json['user']),
      createdAt: DateTime.parse(json['created_at']),
      updatedAt: DateTime.parse(json['updated_at']),
      deletedAt: json['deleted_at'] != null
          ? DateTime.parse(json['deleted_at'])
          : null,
    );
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = {
      'transaction_id': transactionId,
      'user': user.toJson(),
      'created_at': createdAt.toIso8601String(),
      'updated_at': updatedAt.toIso8601String(),
      'deleted_at': deletedAt?.toIso8601String(),
    };
    return data;
  }
}
