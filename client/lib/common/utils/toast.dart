import 'package:flutter_easyloading/flutter_easyloading.dart';

class Toast {
  Toast._();

  static show(String message) {
    EasyLoading.showToast(message);
  }

  static showTODO({String? message}) {
    EasyLoading.showToast(message ?? '功能开发中，敬请期待');
  }
}
