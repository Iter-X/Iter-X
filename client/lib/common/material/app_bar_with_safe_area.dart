import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class AppBarWithSafeArea extends StatelessWidget {
  final Widget child;
  final Color? topColor;
  final Color? bottomColor;
  final bool top;
  final bool bottom;
  final bool left;
  final bool right;
  final Brightness? statusBarIconBrightness;
  final Brightness? statusBarBrightness;
  final bool hasAppBar;
  final String? title;
  final Widget? leading;
  final List<Widget>? actions;
  final Color? backgroundColor;
  final double? elevation;
  final PreferredSizeWidget? appBarBottom;
  final double? toolbarHeight;
  final bool centerTitle;

  const AppBarWithSafeArea({
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
    this.hasAppBar = true,
    this.title,
    this.leading,
    this.actions,
    this.backgroundColor,
    this.elevation,
    this.appBarBottom,
    this.toolbarHeight,
    this.centerTitle = true,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: backgroundColor ?? Colors.transparent,
      appBar: hasAppBar
          ? AppBar(
              backgroundColor: backgroundColor ?? Colors.transparent,
              elevation: elevation ?? 0,
              systemOverlayStyle: SystemUiOverlayStyle(
                statusBarColor: Colors.transparent,
                statusBarIconBrightness:
                    statusBarIconBrightness ?? Brightness.dark,
                statusBarBrightness: statusBarBrightness ?? Brightness.light,
              ),
              leading: leading,
              title: title != null
                  ? Text(
                      style: TextStyle(
                        color: AppColor.primaryFont,
                        fontSize: 18.sp,
                        fontWeight: AppFontWeight.medium,
                      ),
                      title!)
                  : null,
              centerTitle: centerTitle,
              actions: actions,
              bottom: appBarBottom,
              toolbarHeight: toolbarHeight,
            )
          : null,
      body: SafeAreaX(
        topColor: topColor,
        bottomColor: bottomColor,
        top: top,
        bottom: bottom,
        left: left,
        right: right,
        statusBarIconBrightness: statusBarIconBrightness,
        statusBarBrightness: statusBarBrightness,
        child: child,
      ),
    );
  }
}

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
      statusBarIconBrightness:
          widget.statusBarIconBrightness ?? Brightness.dark,
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
