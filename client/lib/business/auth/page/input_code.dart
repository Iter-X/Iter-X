import 'dart:async';

import 'package:client/app/routes.dart';
import 'package:client/business/auth/entity/user_info_entity.dart';
import 'package:client/common/material/loading.dart';
import 'package:client/common/utils/color.dart';
import 'package:client/common/utils/shared_preference_util.dart';
import 'package:client/common/utils/toast.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:pin_code_fields/pin_code_fields.dart';
import 'package:gap/gap.dart';

import '../../../common/material/state.dart';
import '../service/auth_service.dart';

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
  }

  @override
  void dispose() {
    cancelTimer();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
      body: SizedBox(
        width: double.infinity,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              '验证码已发送至',
              style: TextStyle(
                fontSize: 16.sp,
                color: BaseColor.c_1D1F1E,
              ),
            ),
            Gap(7.h),
            Text(
              '+86 ${widget.argument.phone}',
              style: TextStyle(
                fontSize: 30.sp,
                color: BaseColor.c_1D1F1E,
                fontWeight: FontWeight.w600,
              ),
            ),
            Container(
              margin: EdgeInsets.only(
                left: 70.w,
                right: 70.w,
                top: 26.h,
                bottom: 17.h,
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
                  activeFillColor: BaseColor.c_F2F2F2,
                  selectedColor: BaseColor.c_1D1F1E,
                  inactiveColor: BaseColor.c_1D1F1E,
                  activeColor: BaseColor.c_1D1F1E,
                ),
                animationDuration: Duration(milliseconds: 300),
                enableActiveFill: false,
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
                  fontWeight: timeStr == '重新发送' ? FontWeight.w600 : FontWeight.w400,
                ),
              ),
            ),
            GestureDetector(
              onTap: () {
                verifyLogin();
              },
              child: Container(
                width: double.infinity,
                height: 52.h,
                margin: EdgeInsets.only(
                  top: 32.h,
                  left: 73.w,
                  right: 73.w,
                ),
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(24.w),
                  color: BaseColor.c_1D1F1E,
                ),
                alignment: Alignment.center,
                child: isLoading
                    ? const LoadingWidget()
                    : Text(
                        '登录',
                        style: TextStyle(
                          color: Colors.white,
                          fontSize: 20.sp,
                        ),
                      ),
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
    UserInfoEntity? result = await AuthService.verifyLogin(
      widget.argument.phone,
      _codeController.text,
    );
    setState(() {
      isLoading = false;
    });
    if (result != null) {
      await BaseSpUtil.setJSON(SpKeys.TOKEN, result.token);
      await BaseSpUtil.setJSON(SpKeys.USER_INFO, result);
      go(Routes.home, clearStack: true);
    }
  }
}
