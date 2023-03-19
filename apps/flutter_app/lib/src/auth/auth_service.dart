class AuthService {
  Future<bool> login(String username, String password) async {
    return false;
  }

  Future<bool> logout() async {
    return false;
  }

  Future<bool> register(String username, String password) async {
    return false;
  }

  Future<bool> isLoggedIn() async => false;

  Future<String> getAccessToken() async => '';
}
