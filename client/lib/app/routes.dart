import 'package:client/business/auth/page/input_code.dart';
import 'package:client/business/auth/page/login.dart';
import 'package:client/business/auth/page/phone_login.dart';
import 'package:client/business/create_trip/page/card_selection.dart';
import 'package:client/business/create_trip/page/create_trip_home.dart';
import 'package:client/business/create_trip/page/creating_trip.dart';
import 'package:client/business/create_trip/page/poi_search.dart';
import 'package:client/business/create_trip/page/select_date.dart';
import 'package:client/business/home_main/page/home_main.dart';
import 'package:client/business/mine/page/mine_page.dart';
import 'package:client/business/trip/page/trip_main_page.dart';
import 'package:fluro/fluro.dart';
import 'package:flutter/material.dart';

// 页面路由路径
class Routes {
  static bool _alreadyDefined = false;

  // 路由路径定义
  static const String root = '/';
  static const String login = '/auth/login';
  static const String phoneLogin = '/auth/phone_login';
  static const String inputCode = '/auth/input_code';
  static const String homeMain = '/home_main/home_main';
  static const String createTripHome = '/create_trip/create_trip_home';
  static const String cardSelection = '/create_trip/card_selection';
  static const String selectDate = '/create_trip/select_date';
  static const String profile = '/mine/profile';
  static const String poiSearch = '/create_trip/poi_search';
  static const String tripOverview = '/trip/overview';
  static const String creatingTrip = '/create_trip/creating_trip';

  // 无需登录即可访问的页面
  static final List<String> _routesWithoutLogin = [
    login,
    phoneLogin,
    inputCode,
  ];

  // 路由配置映射表
  static final Map<String, Handler> _routeHandlers = {
    root: Handler(handlerFunc: (context, params) => const HomeMainPage()),
    login: Handler(handlerFunc: (context, params) => const LoginPage()),
    phoneLogin:
        Handler(handlerFunc: (context, params) => const PhoneLoginPage()),
    inputCode: Handler(handlerFunc: (context, params) {
      if (context?.settings?.arguments is! InputCodeArgument) {
        return const LoginPage(); // 如果参数不正确，返回登录页
      }
      final args = context!.settings!.arguments as InputCodeArgument;
      return InputCodePage(argument: args);
    }),
    homeMain: Handler(handlerFunc: (context, params) => const HomeMainPage()),
    createTripHome:
        Handler(handlerFunc: (context, params) => const CreateTripHomePage()),
    cardSelection:
        Handler(handlerFunc: (context, params) => const CardSelectionPage()),
    profile: Handler(handlerFunc: (context, params) => const MinePage()),
    poiSearch: Handler(handlerFunc: (context, params) => const PoiSearchPage()),
    tripOverview:
        Handler(handlerFunc: (context, params) => const TripMainPage()),
    selectDate:
        Handler(handlerFunc: (context, params) => const SelectDatePage()),
    creatingTrip: Handler(handlerFunc: (context, params) {
      final args = context!.settings!.arguments as Map<String, dynamic>;
      return CreatingTripPage(params: args['params']);
    }),
  };

  // 404路由处理器
  static final Handler _notFoundHandler = Handler(
    handlerFunc: (context, params) => Scaffold(
      body: Center(
        child: Text('Route not found: ${context?.settings?.name}'),
      ),
    ),
  );

  static void config(FluroRouter router) {
    if (_alreadyDefined) return;
    _alreadyDefined = true;

    // 注册所有路由
    _routeHandlers.forEach((path, handler) {
      router.define(path, handler: handler);
    });

    // 配置404路由
    router.notFoundHandler = _notFoundHandler;
  }

  // 检查路由是否需要登录
  static bool requiresLogin(String? path) {
    return !_routesWithoutLogin.contains(path);
  }
}
