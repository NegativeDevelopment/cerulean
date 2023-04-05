class GroupUserItem {
  final String id;
  final String name;

  GroupUserItem({
    required this.id,
    required this.name,
  });

  factory GroupUserItem.fromJson(Map<String, dynamic> json) {
    return GroupUserItem(
      id: json['id'],
      name: json['name'],
    );
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['id'] = id;
    data['name'] = name;
    return data;
  }
}
