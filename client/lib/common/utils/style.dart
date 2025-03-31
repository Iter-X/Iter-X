import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class BaseStyle {
  // app bar start
  static AppBarTheme appBarTheme = AppBarTheme(
    elevation: 0,
    centerTitle: true,
    titleTextStyle: TextStyle(color: Colors.black),
    iconTheme: const IconThemeData(color: Colors.black),
  );

  // button start
  static ButtonStyle elevatedButtonStyle = ElevatedButton.styleFrom(
    elevation: 0,
    backgroundColor: AppColor.theme,
    minimumSize: const Size(0, 0),
    shadowColor: Colors.transparent,
    splashFactory: NoSplash.splashFactory,
    tapTargetSize: MaterialTapTargetSize.shrinkWrap,
    padding: const EdgeInsets.symmetric(vertical: 8),
    textStyle: TextStyle(fontSize: 16.sp, color: Colors.black),
    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8.w)),
  );

  static ButtonStyle outlinedButtonStyle = OutlinedButton.styleFrom(
    elevation: 0,
    backgroundColor: AppColor.theme,
    minimumSize: const Size(0, 0),
    shadowColor: Colors.transparent,
    splashFactory: NoSplash.splashFactory,
    tapTargetSize: MaterialTapTargetSize.shrinkWrap,
    padding: const EdgeInsets.symmetric(vertical: 8),
    textStyle: TextStyle(fontSize: 16.sp, color: Colors.black),
    shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8.w)),
    side: const BorderSide(width: 1, color: AppColor.theme),
  );
}
