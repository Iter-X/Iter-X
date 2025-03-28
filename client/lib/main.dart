import 'package:client/app/constants.dart';
import 'package:client/common/material/theme_data.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';

import 'app/events/events.dart';
import 'app/notifier/user.dart';
import 'app/routes.dart';
import 'common/material/app.dart';
import 'common/material/state.dart';
import 'common/utils/api_util.dart';

final GlobalKey<NavigatorState> navigatorKey = GlobalKey<NavigatorState>();

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  // 强制竖屏
  SystemChrome.setPreferredOrientations([
    DeviceOrientation.portraitUp,
    DeviceOrientation.portraitDown,
  ]);

  final userNotifier = UserNotifier();
  await userNotifier.loadUserInfo();

  ApiUtil.apiModel = await ApiUtil.getSelectedApiModel();

  // 监听登录过期事件
  eventBus.on<EventUnauthorized>().listen((event) async {
    // 清除用户信息
    await userNotifier.logout();
    // 跳转到登录页
    navigatorKey.currentState?.pushNamedAndRemoveUntil(
      Routes.login,
      (Route<dynamic> route) => false, // 清空所有页面
    );
  });

  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => userNotifier),
      ],
      child: const MyApp(),
    ),
  );
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends BaseState<MyApp> {
  @override
  Widget build(BuildContext context) {
    return BaseApp.create(
      context: context,
      title: Constants.appName,
      configTheme: () {
        return BaseThemeData.create(fontFamily: 'AlibabaPuHuiTi3');
      },
      configRouter: (router) {
        Routes.config(router);
      },
      navigatorKey: navigatorKey,
    );
  }
}
