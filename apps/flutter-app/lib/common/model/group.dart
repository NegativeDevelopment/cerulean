class Group {
  final String name;
  final int memberCount;
  final int transactionCount;

  Group({
    required this.name,
    required this.memberCount,
    required this.transactionCount,
  });

  factory Group.fromJson(Map<String, dynamic> json) {
    return Group(
      name: json['name'],
      memberCount: json['memberCount'],
      transactionCount: json['transactionCount'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'name': name,
      'memberCount': memberCount,
      'transactionCount': transactionCount,
    };
  }
}
