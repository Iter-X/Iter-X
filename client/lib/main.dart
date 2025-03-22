import 'package:client/business/auth/page/input_code.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';
// import 'package:flutter_ali_auth/flutter_ali_auth.dart';

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
  // 阿里云一键登录初始化
  // await AliAuthClient.initSdk(
  //   authConfig: AuthConfig(
  //     iosSdk: Constants.aliIosSdk,
  //     androidSdk: Constants.aliAndroidSdk,
  //     enableLog: false,
  //     authUIConfig: FullScreenUIConfig(
  //       backgroundColor: '#f2f2f2',
  //       phoneNumberConfig: PhoneNumberConfig(
  //         numberColor: '#1D1F1E',
  //         numberFontSize: 30,
  //       ),
  //       loginButtonConfig: LoginButtonConfig(
  //         loginBtnText: '一键登录',
  //         loginBtnTextColor: '#ffffff',
  //         loginBtnTextSize: 16,
  //         loginBtnHeight: 52,
  //         loginBtnWidth: 285,
  //         loginBtnNormalImage: 'assets/images/login_btn_normal.png',
  //         loginBtnFrameOffsetY: 60,
  //       ),
  //       changeButtonConfig: ChangeButtonConfig(
  //         changeBtnTitle: '手机验证码登录',
  //         changeBtnTextColor: '#1D1F1E',
  //         changeBtnTextSize: 16,
  //       ),
  //       // customViewBlockList:
  //     ),
  //   ),
  // );
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
