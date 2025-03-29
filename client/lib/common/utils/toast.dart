import 'package:client/common/utils/color.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:fluttertoast/fluttertoast.dart';

class ToastX {
  static void show(String message, {BuildContext? context}) {
    Fluttertoast.showToast(
      msg: message,
      toastLength: Toast.LENGTH_SHORT,
      gravity: ToastGravity.TOP,
      timeInSecForIosWeb: 1,
      backgroundColor: BaseColor.primary.withOpacity(0.8),
      textColor: BaseColor.secondary,
      fontSize: 16.sp,
    );
  }

  static void showTODO({String? message, BuildContext? context}) {
    show(message ?? '功能开发中，敬请期待', context: context);
  }
}
