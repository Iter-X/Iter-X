import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';

import 'app/constants.dart';
import 'app/notifier/user.dart';
import 'app/routes.dart';
import 'business/auth/entity/user_info_entity.dart';
import 'common/material/app.dart';
import 'common/material/state.dart';
import 'common/material/theme_data.dart';
import 'common/utils/api_util.dart';
import 'common/utils/logger.dart';
import 'common/utils/shared_preference_util.dart';

void main() async {
  // 强制竖屏
  WidgetsFlutterBinding.ensureInitialized();
  SystemChrome.setPreferredOrientations([
    DeviceOrientation.portraitUp,
    DeviceOrientation.portraitDown,
  ]);
  ApiUtil.apiModel = await ApiUtil.getSelectedApiModel();
  UserInfoEntity? user;
  if (await BaseSpUtil.getJSON(SpKeys.USER_INFO) != null) {
    user = UserInfoEntity.fromJson(await BaseSpUtil.getJSON(SpKeys.USER_INFO));
    BaseLogger.v('load user from sp: ${user.toJson()}');
    Routes.needLogin = false;
  }
  runApp(
    MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (_) => UserNotifier(user)),
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
        return BaseThemeData.create();
      },
      configRouter: (router) {
        Routes.config(router);
      },
    );
  }
}
