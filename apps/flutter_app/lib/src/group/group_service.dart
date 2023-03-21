class GroupService {
  final String _baseUrl = 'https://api.groupme.com/v3';
  final String _token;

  GroupService(this._token);

  // Future<List<Group>> getGroups() async {
  //   final response = await http.get('$_baseUrl/groups?token=$_token');
  //   if (response.statusCode == 200) {
  //     final groups = json.decode(response.body)['response'];
  //     return groups.map<Group>((group) => Group.fromJson(group)).toList();
  //   } else {
  //     throw Exception('Failed to load groups');
  //   }
  // }
}
