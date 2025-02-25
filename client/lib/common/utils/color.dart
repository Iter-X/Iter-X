import 'package:flutter/material.dart';

// 颜色工具类
class BaseColor {
  //
  static MaterialColor themeMaterialColor = MaterialColor(
    _themePrimaryValue,
    <int, Color>{
      50: const Color(_themePrimaryValue).withOpacity(0.05),
      100: const Color(_themePrimaryValue).withOpacity(0.1),
      200: const Color(_themePrimaryValue).withOpacity(0.2),
      300: const Color(_themePrimaryValue).withOpacity(0.3),
      400: const Color(_themePrimaryValue).withOpacity(0.4),
      500: const Color(_themePrimaryValue),
      600: const Color(_themePrimaryValue).withOpacity(0.6),
      700: const Color(_themePrimaryValue).withOpacity(0.7),
      800: const Color(_themePrimaryValue).withOpacity(0.8),
      900: const Color(_themePrimaryValue).withOpacity(0.9),
    },
  );
  static const int _themePrimaryValue = 0xFFB9F32B;
  static const Color theme = Color(_themePrimaryValue);
  static const Color divider = Color(0xFFE9E9E9);
  static const Color scaffoldBackgroundColor = Color(0xFFF2F2F2);

  // text color
  static const Color title = Color(0xFF000000);
  static const Color content = Color(0xFF111111);
  static const Color hint = Color(0xFF535A5F);

  static const Color c_1D1F1E = Color(0xFF1D1F1E);
  static const Color c_F2F2F2 = Color(0xFFF2F2F2);
}
