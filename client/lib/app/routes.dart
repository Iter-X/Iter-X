import 'package:client/business/auth/page/login.dart';
import 'package:fluro/fluro.dart';

// 页面路由路径
class Routes {
  static bool _hasDefaultRoute = false;

  static String root = '/';

  // 登录
  static String login = '/auth/login';

  static void config(FluroRouter router) {
    if (!_hasDefaultRoute) {
      router.define(root, handler: Handler(handlerFunc: (c, p) {
        return const LoginPage();
      }));
      _hasDefaultRoute = true;
    }
  }
}