import 'package:client/business/auth/page/input_code.dart';
import 'package:client/business/auth/page/login.dart';
import 'package:client/business/auth/page/phone_login.dart';
import 'package:client/business/home/page/home.dart';
import 'package:fluro/fluro.dart';

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
  static String home = '/home/home';

  // 无需登录即可访问的页面
  static List<String> routesWithoutLogin = [
    login,
    phoneLogin,
    inputCode,
  ];

  static void config(FluroRouter router) {
    if (!_hasDefaultRoute) {
      router.define(root, handler: Handler(handlerFunc: (c, p) {
        return const LoginPage();
      }));
      _hasDefaultRoute = true;
    }
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
    router.define(home, handler: Handler(handlerFunc: (c, p) {
      return const HomePage();
    }));
  }
}