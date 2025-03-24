import 'package:fluro/fluro.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

import 'over_scroll.dart';

class SafeAreaX extends StatefulWidget {
  final Widget child;
  final Color? topColor;
  final Color? bottomColor;
  final bool top;
  final bool bottom;
  final bool left;
  final bool right;
  final Brightness? statusBarIconBrightness;
  final Brightness? statusBarBrightness;

  const SafeAreaX({
    super.key,
    required this.child,
    this.topColor,
    this.bottomColor,
    this.top = true,
    this.bottom = true,
    this.left = true,
    this.right = true,
    this.statusBarIconBrightness,
    this.statusBarBrightness,
  });

  @override
  State<SafeAreaX> createState() => _SafeAreaXState();
}

class _SafeAreaXState extends State<SafeAreaX> {
  @override
  void initState() {
    super.initState();
    _updateStatusBarStyle();
  }

  @override
  void didUpdateWidget(SafeAreaX oldWidget) {
    super.didUpdateWidget(oldWidget);
    _updateStatusBarStyle();
  }

  void _updateStatusBarStyle() {
    SystemChrome.setSystemUIOverlayStyle(SystemUiOverlayStyle(
      statusBarColor: Colors.transparent,
      statusBarIconBrightness: widget.statusBarIconBrightness ?? Brightness.dark,
      statusBarBrightness: widget.statusBarBrightness ?? Brightness.light,
    ));
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        if (widget.top)
          Container(
            color: widget.topColor ?? Colors.transparent,
            child: SafeArea(
              bottom: false,
              left: widget.left,
              right: widget.right,
              child: Container(),
            ),
          ),
        Expanded(
          child: SafeArea(
            top: false,
            bottom: false,
            left: widget.left,
            right: widget.right,
            child: widget.child,
          ),
        ),
        if (widget.bottom)
          Container(
            color: widget.bottomColor ?? Colors.transparent,
            child: SafeArea(
              top: false,
              left: widget.left,
              right: widget.right,
              child: Container(),
            ),
          ),
      ],
    );
  }
}

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
    initialRoute,
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
    // 配置页面路由
    if (onGenerateRoute == null) {
      assert(configRouter != null);
      configRouter?.call(router);
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
        return MaterialApp(
          key: key,
          navigatorKey: navigatorKey,
          scaffoldMessengerKey: scaffoldMessengerKey,
          home: home,
          routes: routes ?? const <String, WidgetBuilder>{},
          initialRoute: initialRoute,
          onGenerateRoute: onGenerateRoute ?? router.generator,
          onGenerateInitialRoutes: onGenerateInitialRoutes,
          onUnknownRoute: onUnknownRoute,
          navigatorObservers: navigatorObservers ?? const <NavigatorObserver>[],
          builder: EasyLoading.init(
            builder: (context, child) => Scaffold(
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
                    child: child!,
                  ),
                ),
              ),
            ),
          ),
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
