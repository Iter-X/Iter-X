import 'package:client/app/constants.dart';
import 'package:client/common/utils/style.dart';
import 'package:flutter/material.dart';

class BaseThemeData {
  BaseThemeData._();

  static ThemeData create({String? fontFamily}) {
    var theme = ThemeData();
    TextSelectionThemeData textSelectionTheme =
        const TextSelectionThemeData().copyWith(
      cursorColor: AppColor.hint, // 输入框 光标颜色
    );
    CardTheme cardTheme = const CardTheme().copyWith(
      elevation: 0,
      color: Colors.white,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(8),
      ),
    );
    ElevatedButtonThemeData elevatedButtonTheme = ElevatedButtonThemeData(
      style: BaseStyle.elevatedButtonStyle,
    );
    TextButtonThemeData textButtonTheme = TextButtonThemeData(
      style: TextButton.styleFrom(
        elevation: 0,
        backgroundColor: AppColor.theme,
        minimumSize: const Size(0, 0),
        shadowColor: Colors.transparent,
        splashFactory: NoSplash.splashFactory,
        tapTargetSize: MaterialTapTargetSize.shrinkWrap,
        padding: const EdgeInsets.symmetric(vertical: 8),
      ),
    );
    OutlinedButtonThemeData outlinedButtonTheme = OutlinedButtonThemeData(
      style: BaseStyle.outlinedButtonStyle,
    );
    DividerThemeData dividerTheme = theme.dividerTheme.copyWith(
      color: AppColor.divider,
      space: 0.5,
      thickness: 0.5,
    );
    return ThemeData(
      primarySwatch: AppColor.themeMaterialColor,
      primaryColor: AppColor.theme,
      primaryColorLight: AppColor.theme,
      primaryColorDark: AppColor.theme,
      scaffoldBackgroundColor: AppColor.scaffoldBackgroundColor,
      appBarTheme: BaseStyle.appBarTheme,
      cardTheme: cardTheme,
      elevatedButtonTheme: elevatedButtonTheme,
      outlinedButtonTheme: outlinedButtonTheme,
      textButtonTheme: textButtonTheme,
      textSelectionTheme: textSelectionTheme,
      dividerTheme: dividerTheme,
      fontFamily: fontFamily,
    );
  }
}
