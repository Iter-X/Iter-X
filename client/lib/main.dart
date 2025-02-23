import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';

import 'app/constants.dart';
import 'app/routes.dart';
import 'common/material/app.dart';
import 'common/material/state.dart';
import 'common/material/theme_data.dart';

void main() async {
  // 强制竖屏
  WidgetsFlutterBinding.ensureInitialized();
  SystemChrome.setPreferredOrientations([
    DeviceOrientation.portraitUp,
    DeviceOrientation.portraitDown,
  ]);
  runApp(
    const MyApp(),
    // MultiProvider(
    //   providers: [],
    //   child: const MyApp(),
    // ),
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
