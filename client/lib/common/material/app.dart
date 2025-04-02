import 'package:client/app/notifier/user.dart';
import 'package:client/app/routes.dart';
import 'package:fluro/fluro.dart';
import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:provider/provider.dart';

import 'over_scroll.dart';

final router = FluroRouter();

typedef ConfigRouter = Function(FluroRouter router);
typedef ConfigTheme = ThemeData Function();

class BaseApp {
  BaseApp._();

  static Widget create({
    Key? key,
    BuildContext? context,
    ConfigRouter? configRouter,
    ConfigTheme? configTheme,
    navigatorKey,
    scaffoldMessengerKey,
    home,
    Map<String, WidgetBuilder>? routes,
    String? initialRoute,
    onGenerateRoute,
    onGenerateInitialRoutes,
    onUnknownRoute,
    List<NavigatorObserver>? navigatorObservers,
    builder,
    title = '',
    onGenerateTitle,
    color,
    theme,
    darkTheme,
    highContrastTheme,
    highContrastDarkTheme,
    themeMode,
    locale,
    localizationsDelegates,
    localeListResolutionCallback,
    localeResolutionCallback,
    supportedLocales,
    debugShowMaterialGrid,
    showPerformanceOverlay,
    checkerboardRasterCacheImages,
    checkerboardOffscreenLayers,
    showSemanticsDebugger,
    debugShowCheckedModeBanner,
    shortcuts,
    actions,
    restorationScopeId,
    scrollBehavior,
    useInheritedMediaQuery,
  }) {
    // confirm the router is configured
    if (configRouter != null) {
      configRouter(router);
    }

    const designSize = Size(430, 932);

    return ScreenUtilInit(
      designSize: designSize,
      fontSizeResolver: (fontSize, instance) {
        final display = View.of(context!).display;
        final screenSize = display.size / display.devicePixelRatio;
        final scaleWidth = screenSize.width / designSize.width;
        return fontSize * scaleWidth;
      },
      minTextAdapt: true,
      splitScreenMode: true,
      builder: (BuildContext context, Widget? child) {
        // check login status
        final userNotifier = context.watch<UserNotifier>();
        // 由于 MaterialApp 的 builder 不能是 async 的，我们需要先用 isTokenExpired 做初步判断
        // 后续的路由判断会使用 checkTokenExpired 进行更准确的检查
        final effectiveInitialRoute = initialRoute ??
            (userNotifier.isTokenExpired ? Routes.login : Routes.homeMain);

        return MaterialApp(
          key: key,
          navigatorKey: navigatorKey,
          scaffoldMessengerKey: scaffoldMessengerKey,
          home: null,
          // use routes instead of home
          routes: const <String, WidgetBuilder>{},
          initialRoute: effectiveInitialRoute,
          onGenerateRoute: (settings) {
            // 如果路由需要登录，先用 isTokenExpired 做快速检查
            // 实际跳转时会通过 go 方法进行更严格的 checkTokenExpired 检查
            if (!Routes.requiresLogin(settings.name) &&
                userNotifier.isTokenExpired) {
              return router.generator(RouteSettings(name: Routes.login));
            }
            return router.generator(settings);
          },
          onGenerateInitialRoutes: (String initialRoute) {
            return [
              router.generator(RouteSettings(name: initialRoute)),
            ].whereType<Route<dynamic>>().toList();
          },
          onUnknownRoute: onUnknownRoute,
          navigatorObservers: navigatorObservers ?? const <NavigatorObserver>[],
          builder: (context, child) {
            if (child == null) return const SizedBox.shrink();

            return Scaffold(
              resizeToAvoidBottomInset: false,
              body: GestureDetector(
                onTap: () {
                  // 页面点击关闭键盘
                  FocusScopeNode currentFocus = FocusScope.of(context);
                  if (!currentFocus.hasPrimaryFocus &&
                      currentFocus.focusedChild != null) {
                    FocusManager.instance.primaryFocus?.unfocus();
                  }
                },
                child: ScrollConfiguration(
                  behavior: OverScrollBehavior(), // 取消滚动组件滑到顶部和尾部水波纹效果
                  child: MediaQuery(
                    data: MediaQuery.of(context).copyWith(textScaleFactor: 1.0),
                    child: child,
                  ),
                ),
              ),
            );
          },
          title: title ?? '',
          onGenerateTitle: onGenerateTitle,
          color: color,
          theme: theme ?? configTheme?.call(),
          darkTheme: darkTheme,
          highContrastTheme: highContrastTheme,
          highContrastDarkTheme: highContrastDarkTheme,
          themeMode: themeMode ?? ThemeMode.system,
          locale: locale,
          localizationsDelegates: localizationsDelegates ??
              [
                GlobalMaterialLocalizations.delegate,
                GlobalWidgetsLocalizations.delegate,
                GlobalCupertinoLocalizations.delegate,
              ],
          localeListResolutionCallback: localeListResolutionCallback,
          localeResolutionCallback: localeResolutionCallback,
          supportedLocales:
              supportedLocales ?? const <Locale>[Locale('zh', 'CH')],
          debugShowMaterialGrid: debugShowMaterialGrid ?? false,
          showPerformanceOverlay: showPerformanceOverlay ?? false,
          checkerboardRasterCacheImages: checkerboardRasterCacheImages ?? false,
          checkerboardOffscreenLayers: checkerboardOffscreenLayers ?? false,
          showSemanticsDebugger: showSemanticsDebugger ?? false,
          debugShowCheckedModeBanner: debugShowCheckedModeBanner ?? false,
          shortcuts: shortcuts,
          actions: actions,
          restorationScopeId: restorationScopeId,
          scrollBehavior: scrollBehavior,
          useInheritedMediaQuery: useInheritedMediaQuery ?? false,
        );
      },
    );
  }
}
