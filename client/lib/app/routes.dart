/*
 * @Description: 
 * @Version: 0.1
 * @Autor: GiottoLLL7
 * @Date: 2025-03-14 13:13:11
 * @LastEditors: GiottoLLL7
 * @LastEditTime: 2025-03-18 23:13:03
 */
import 'package:client/business/auth/page/input_code.dart';
import 'package:client/business/auth/page/login.dart';
import 'package:client/business/auth/page/phone_login.dart';
import 'package:client/business/home/page/home.dart';
import 'package:client/business/create_trip/page/card_selection.dart';
import 'package:client/business/create_trip/page/create_trip_home.dart';
import 'package:fluro/fluro.dart';

import '../business/home_main/page/home_main.dart';

// 页面路由路径
class Routes {
  static bool _hasDefaultRoute = false;
  static bool needLogin = false;

  static String root = '/';

  // 登录
  static String login = '/auth/login';
  // 手机号登录
  static String phoneLogin = '/auth/phone_login';
  // 输入验证码页面
  static String inputCode = '/auth/input_code';
  // 首页
  static String homeMain = '/home_main/home_main';
  // 创建行程
  static String createTripHome = '/create_trip/create_trip_home';
  // 创建行程图卡选择
  static String cardSelection = '/create_trip/card_selection';

  // 无需登录即可访问的页面
  static List<String> routesWithoutLogin = [
    login,
    phoneLogin,
    inputCode,
    createTripHome,
  ];

  static void config(FluroRouter router) {
    if (!_hasDefaultRoute) {
      router.define(root, handler: Handler(handlerFunc: (c, p) {
        return const HomeMainPage();
      }));
      _hasDefaultRoute = true;
    }
    router.define(login, handler: Handler(handlerFunc: (c, p) {
      return const CardSelectionPage();
    }));
    router.define(login, handler: Handler(handlerFunc: (c, p) {
      return const LoginPage();
    }));
    router.define(phoneLogin, handler: Handler(handlerFunc: (c, p) {
      return const PhoneLoginPage();
    }));
    router.define(inputCode, handler: Handler(handlerFunc: (c, p) {
      final args = c?.settings?.arguments as InputCodeArgument;
      return InputCodePage(argument: args);
    }));
    router.define(homeMain, handler: Handler(handlerFunc: (c, p) {
      return const HomeMainPage();
    }));
    router.define(createTripHome, handler: Handler(handlerFunc: (c, p) {
      return const CreateTripHomePage();
    }));
    router.define(cardSelection, handler: Handler(handlerFunc: (c, p) {
      return const CardSelectionPage();
    }));
  }
}
