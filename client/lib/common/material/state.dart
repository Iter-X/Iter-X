import 'dart:async';

import 'package:connectivity_plus/connectivity_plus.dart';
import 'package:client/common/utils/color.dart';
import 'package:easy_refresh/easy_refresh.dart';
import 'package:event_bus/event_bus.dart';
import 'package:fluro/fluro.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../../app/notifier/user.dart';
import '../../app/routes.dart';
import '../utils/logger.dart';
import 'app.dart';

/// The global [EventBus] object.
EventBus eventBus = EventBus();

/// 事件通知
abstract class Event {
  /// 监听事件
  void onEvent(void Function(dynamic event) onData);

  /// 通知事件
  void fireEvent(event);
}

/*
所有控件基础state
 */
abstract class BaseState<T extends StatefulWidget> extends State<T> with Event {
  //
  StreamSubscription? _eventSub;

  //
  List<ConnectivityResult>? _connectivity;

  //
  StreamSubscription? _connectSub;

  //
  // ScaffoldState _scaffoldState = ScaffoldState.none;

  bool isShowNewMessage = false; // 是否展示新消息弹窗

  @override
  void initState() {
    super.initState();
    //
    if (checkConnectivity) {
      _connectSub = Connectivity().onConnectivityChanged.listen((List<ConnectivityResult> result) {
        refreshConnectivityResult();
      });
    }

  }

  @override
  void dispose() {
    _connectSub?.cancel();
    _eventSub?.cancel();
    super.dispose();
  }

  @override
  void onEvent(void Function(dynamic event) onData) {
    _eventSub ??= eventBus.on().listen((event) {
      if (mounted) {
        onData(event);
      }
    });
  }

  @override
  void fireEvent(event) {
    eventBus.fire(event);
  }

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        buildContent(),
      ],
    );
  }

  Widget buildContent() {
    return Container();
  }

  ///
  bool get checkConnectivity => false;

  void refreshConnectivityResult() async {
    var result = await (Connectivity().checkConnectivity());
    setState(() {
      _connectivity = result;
    });
  }

  Future go(
    String path, {
    bool? needLogin,
    arguments,
    bool replace = false,
    bool clearStack = false,
    TransitionType? transition,
  }) {
    BaseLogger.v('go path: $path');
    // 关闭键盘
    FocusScope.of(context).unfocus();
    //
    needLogin ??= !Routes.routesWithoutLogin.contains(path);
    // //
    if (needLogin && !context.read<UserNotifier>().logined()) {
      return go(Routes.phoneLogin).then((value) {
        if (context.read<UserNotifier>().logined()) {
          return router.navigateTo(
            context, path, //路径
            replace: replace,
            clearStack: clearStack,
            transition: transition ?? TransitionType.cupertino,
            routeSettings: RouteSettings(
              arguments: arguments,
            ),
          );
        }
      });
    } else {
      return router.navigateTo(
        context,
        path,
        replace: replace,
        clearStack: clearStack,
        transition: transition ?? TransitionType.cupertino,
        routeSettings: RouteSettings(
          arguments: arguments,
        ),
      );
    }
    return Future(() => null);
  }
}

buildHeader({
  TextStyle? textStyle,
  TextStyle? messageStyle,
  IconThemeData? iconTheme,
}) {
  return ClassicHeader(
    dragText: '下拉可以刷新',
    armedText: '松开立即刷新',
    readyText: '正在刷新数据中...',
    processedText: '刷新成功',
    failedText: '刷新失败',
    noMoreText: '没有更多数据',
    showText: true,
    messageText: '最后更新：今天 %T',
    textStyle: textStyle,
    messageStyle: messageStyle,
    iconTheme: iconTheme,
  );
}

buildFooter({String? loadedText}) {
  return ClassicFooter(
    // loadText: 'loadText',
    // loadReadyText: 'loadReadyText',
    processingText: '正在刷新',
    processedText: loadedText ?? '加载成功',
    failedText: '加载失败',
    noMoreText: '没有更多数据',
    showText: false,
    textStyle: TextStyle(color: BaseColor.hint),
  );
}
