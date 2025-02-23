import 'package:flutter/material.dart';

/*
* $作用：去除ListView 滑動波紋
* $參考：https://blog.csdn.net/u013894711/article/details/102572581
* */
class OverScrollBehavior extends ScrollBehavior {
  @override
  Widget buildViewportChrome(
      BuildContext context, Widget child, AxisDirection axisDirection) {
    switch (getPlatform(context)) {
      case TargetPlatform.iOS:
        return child;
      case TargetPlatform.android:
      case TargetPlatform.fuchsia:
        return GlowingOverscrollIndicator(
          showLeading: false,
          // 不显示尾部水波纹
          showTrailing: false,
          axisDirection: axisDirection,
          color: Colors.transparent,
          child: child,
        );
      case TargetPlatform.linux:
        return child;
      case TargetPlatform.macOS:
        return child;
      case TargetPlatform.windows:
        return child;
    }
  }
}
