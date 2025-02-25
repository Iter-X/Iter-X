import 'package:client/app/routes.dart';
import 'package:client/business/auth/page/input_code.dart';
import 'package:client/common/material/text_field.dart';
import 'package:client/common/utils/color.dart';
import 'package:client/common/utils/toast.dart';
import 'package:client/common/utils/util.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

import '../../../common/material/state.dart';

class PhoneLoginPage extends StatefulWidget {
  const PhoneLoginPage({super.key});

  @override
  State<PhoneLoginPage> createState() => _PhoneLoginPageState();
}

class _PhoneLoginPageState extends BaseState<PhoneLoginPage> {
  late TextEditingController _controller;

  @override
  void initState() {
    _controller = TextEditingController();
    super.initState();
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Container(
        margin: EdgeInsets.symmetric(horizontal: 73.w),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              '请输入手机号',
              style: TextStyle(
                color: BaseColor.c_1D1F1E,
                fontSize: 30.w,
                fontWeight: FontWeight.w600,
              ),
            ),
            Gap(6.h),
            Text(
              '首次登录自动创建账号',
              style: TextStyle(
                color: BaseColor.c_1D1F1E,
                fontSize: 16.w,
              ),
            ),
            Container(
              width: double.infinity,
              height: 52.h,
              margin: EdgeInsets.only(top: 40.h),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(30.w),
                border: Border.all(
                  color: BaseColor.c_1D1F1E,
                  width: 1.w,
                ),
              ),
              padding: EdgeInsets.only(left: 23.w),
              child: Row(
                children: [
                  Text(
                    '+86',
                    style: TextStyle(
                      color: BaseColor.c_1D1F1E,
                      fontSize: 20.sp,
                    ),
                  ),
                  Expanded(
                    child: BaseTextFieldWidget(
                      controller: _controller,
                      contentPadding: EdgeInsets.only(
                        left: 17.w,
                        right: 17.w,
                      ),
                      hintText: '输入手机号',
                      hintStyle: TextStyle(
                        color: BaseColor.hint,
                        fontSize: 20.sp,
                        fontWeight: FontWeight.w400,
                      ),
                      style: TextStyle(
                        color: BaseColor.c_1D1F1E,
                        fontSize: 20.sp,
                        fontWeight: FontWeight.w600,
                      ),
                      lengthLimit: 11,
                    ),
                  ),
                ],
              ),
            ),
            GestureDetector(
              onTap: () => codeLogin(),
              child: Container(
                width: double.infinity,
                height: 52.h,
                margin: EdgeInsets.only(top: 40.h),
                decoration: BoxDecoration(
                  borderRadius: BorderRadius.circular(24.w),
                  color: BaseColor.c_1D1F1E,
                ),
                alignment: Alignment.center,
                child: Text(
                  '发送短信验证码',
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

  void codeLogin() {
    if (BaseUtil.isEmpty(_controller.text)) {
      Toast.show('请输入手机号');
      return;
    }
    if (_controller.text.length != 11) {
      Toast.show('请输入正确的手机号');
      return;
    }
    go(
      Routes.inputCode,
      arguments: InputCodeArgument(phone: _controller.text.trim()),
    );
  }
}
