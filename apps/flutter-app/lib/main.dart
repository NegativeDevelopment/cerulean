import 'package:cerulean_app/home/home_page.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'common/model/config.dart';
import 'common/model/token.dart';
import 'login/login_page.dart';

void main() {
  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (context) => TokenModel()),
        ChangeNotifierProvider(create: (context) => ConfigModel()),
      ],
      child: MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Cerulean App',
      theme: ThemeData.light(useMaterial3: true),
      darkTheme: ThemeData.dark(useMaterial3: true),
      home: _buildPage(context),
    );
  }

  Widget _buildPage(BuildContext context) {
    var tokenModel = context.watch<TokenModel>();

    if (!tokenModel.isLoggedIn()) {
      return LoginPage();
    } else {
      return const HomePage();
    }
  }
}
