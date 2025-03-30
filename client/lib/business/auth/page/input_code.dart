import 'dart:async';

import 'package:client/app/notifier/user.dart';
import 'package:client/app/routes.dart';
import 'package:client/business/auth/service/auth_service.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/loading.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/app_config.dart';
import 'package:client/common/utils/color.dart';
import 'package:client/common/widgets/base_button.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';
import 'package:pin_code_fields/pin_code_fields.dart';
import 'package:provider/provider.dart';

class InputCodeArgument {
  final String phone;

  const InputCodeArgument({
    required this.phone,
  });
}

class InputCodePage extends StatefulWidget {
  final InputCodeArgument argument;

  const InputCodePage({
    super.key,
    required this.argument,
  });

  @override
  State<InputCodePage> createState() => _InputCodePageState();
}

class _InputCodePageState extends BaseState<InputCodePage> {
  late TextEditingController _codeController;
  String timeStr = '重新发送';
  Timer? _timer;
  int time = 60;
  bool isLoading = false;
  final FocusNode _focusNode = FocusNode();

  @override
  void initState() {
    _codeController = TextEditingController();
    _codeController.addListener(() {
      if (_codeController.text.length == 6) {
        verifyLogin();
      }
    });
    super.initState();
    startTimer();
    Future.delayed(const Duration(milliseconds: 100), () {
      if (mounted) {
        _focusNode.requestFocus();
      }
    });
  }

  @override
  void dispose() {
    cancelTimer();
    _focusNode.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      hasAppBar: true,
      backgroundColor: BaseColor.bg,
      leading: ReturnButton(),
      child: SizedBox(
        width: double.infinity,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              '验证码已发送至',
              style: TextStyle(
                fontSize: 16.sp,
                color: BaseColor.primaryFont,
              ),
            ),
            Gap(10.h),
            Text(
              '+86 ${widget.argument.phone}',
              style: TextStyle(
                fontSize: 28.sp,
                color: BaseColor.primaryFont,
                fontWeight: AppFontWeight.medium,
              ),
            ),
            Container(
              margin: EdgeInsets.only(
                left: 70.w,
                right: 70.w,
                top: 20.h,
                bottom: 10.h,
              ),
              child: PinCodeTextField(
                length: 6,
                obscureText: false,
                animationType: AnimationType.fade,
                pinTheme: PinTheme(
                  shape: PinCodeFieldShape.box,
                  borderRadius: BorderRadius.circular(12.w),
                  fieldHeight: 52.h,
                  fieldWidth: 42.w,
                  selectedColor: BaseColor.inputGrayBG,
                  inactiveColor: BaseColor.inputGrayBG,
                  activeColor: BaseColor.inputGrayBG,
                  selectedFillColor: BaseColor.inputGrayBG,
                  activeFillColor: BaseColor.inputGrayBG,
                  inactiveFillColor: BaseColor.inputGrayBG,
                ),
                textStyle: TextStyle(
                  fontSize: 28.sp,
                  color: BaseColor.primaryFont,
                  fontWeight: AppFontWeight.regular,
                ),
                animationDuration: Duration(milliseconds: 300),
                cursorColor: BaseColor.primary,
                cursorWidth: 2,
                cursorHeight: 28.h,
                enableActiveFill: true,
                controller: _codeController,
                onCompleted: (v) {},
                onChanged: (value) {},
                beforeTextPaste: (text) {
                  return true;
                },
                appContext: context,
              ),
            ),
            GestureDetector(
              onTap: () {
                if (timeStr == '重新发送') {
                  startTimer();
                }
              },
              child: Text(
                timeStr,
                style: TextStyle(
                  fontSize: 16.sp,
                  color: BaseColor.c_1D1F1E,
                  fontWeight: timeStr == '重新发送'
                      ? AppFontWeight.medium
                      : AppFontWeight.regular,
                ),
              ),
            ),
            Container(
              margin: EdgeInsets.only(
                top: 20.h,
                left: 72.w,
                right: 72.w,
              ),
              child: isLoading
                  ? const LoadingWidget()
                  : BaseButton(
                      text: '登录',
                      textSize: 18.sp,
                      textColor: Colors.white,
                      backgroundColor: BaseColor.c_1D1F1E,
                      onTap: () => verifyLogin(),
                    ),
            ),
          ],
        ),
      ),
    );
  }

  // 启动定时器
  void startTimer() {
    _timer = Timer.periodic(
      Duration(seconds: 1),
      (timer) {
        if (time == 0) {
          setState(() {
            timeStr = '重新发送';
          });
          time = 60;
          cancelTimer();
        } else {
          time--;
          setState(() {
            timeStr = '${time}s';
          });
        }
      },
    );
  }

  void cancelTimer() {
    _timer?.cancel();
    _timer = null;
  }

  void verifyLogin() async {
    if (isLoading) {
      return;
    }
    setState(() {
      isLoading = true;
    });
    final token = await AuthService.verifyLogin(
      widget.argument.phone,
      _codeController.text,
    );
    setState(() {
      isLoading = false;
    });
    if (token != null) {
      if (mounted) {
        // guard the use of BuildContext with the mounted check
        UserNotifier userNotifier =
            Provider.of<UserNotifier>(context, listen: false);
        await userNotifier.login(token: token);
        await userNotifier.refreshUserInfo();
      }

      go(Routes.homeMain, clearStack: true);
    }
  }
}
