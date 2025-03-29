import 'package:client/app/routes.dart';
import 'package:client/business/auth/page/input_code.dart';
import 'package:client/business/auth/service/auth_service.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/material/loading.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/material/text_field.dart';
import 'package:client/common/utils/color.dart';
import 'package:client/common/utils/toast.dart';
import 'package:client/common/utils/util.dart';
import 'package:client/common/widgets/base_button.dart';
import 'package:client/common/widgets/return_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';

class PhoneLoginPage extends StatefulWidget {
  const PhoneLoginPage({super.key});

  @override
  State<PhoneLoginPage> createState() => _PhoneLoginPageState();
}

class _PhoneLoginPageState extends BaseState<PhoneLoginPage> {
  late TextEditingController _controller;
  bool isLoading = false;
  final FocusNode _focusNode = FocusNode();

  @override
  void initState() {
    _controller = TextEditingController();
    super.initState();

    // get focus after 100ms
    Future.delayed(const Duration(milliseconds: 100), () {
      if (mounted) {
        _focusNode.requestFocus();
      }
    });
  }

  @override
  void dispose() {
    _controller.dispose();
    _focusNode.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      hasAppBar: true,
      backgroundColor: BaseColor.bg,
      leading: ReturnButton(),
      child: Container(
        margin: EdgeInsets.symmetric(horizontal: 72.w),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              '请输入手机号',
              style: TextStyle(
                color: BaseColor.primaryFont,
                fontSize: 28.w,
                fontWeight: FontWeight.w600,
              ),
            ),
            Gap(10.h),
            Text(
              '首次登录自动创建账号',
              style: TextStyle(
                color: BaseColor.primaryFont,
                fontSize: 16.w,
              ),
            ),
            Container(
              width: double.infinity,
              height: 52.h,
              margin: EdgeInsets.only(top: 40.h),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(30.r),
                border: Border.all(
                  color: BaseColor.borderLine,
                  width: 1.w,
                ),
              ),
              padding: EdgeInsets.only(left: 24.w),
              child: Row(
                children: [
                  Text(
                    '+86',
                    style: TextStyle(
                      color: BaseColor.primaryFont,
                      fontSize: 18.sp,
                    ),
                  ),
                  Expanded(
                    child: BaseTextFieldWidget(
                      controller: _controller,
                      focusNode: _focusNode,
                      contentPadding: EdgeInsets.only(
                        left: 24.w,
                        right: 24.w,
                      ),
                      hintText: '输入手机号',
                      hintStyle: TextStyle(
                        color: BaseColor.hint,
                        fontSize: 18.sp,
                        fontWeight: FontWeight.w400,
                      ),
                      style: TextStyle(
                        color: BaseColor.c_1D1F1E,
                        fontSize: 18.sp,
                        fontWeight: FontWeight.w500,
                      ),
                      lengthLimit: 11,
                    ),
                  ),
                ],
              ),
            ),
            Container(
              margin: EdgeInsets.only(top: 40.h),
              child: BaseButton(
                text: '发送短信验证码',
                isLoading: isLoading,
                loadingWidget: const LoadingWidget(
                  color: BaseColor.secondary,
                ),
                textSize: 18.sp,
                textColor: BaseColor.secondary,
                onTap: () => codeLogin(),
              ),
            ),
          ],
        ),
      ),
    );
  }

  Future<void> codeLogin() async {
    if (isLoading) {
      return;
    }
    String phoneNumber = _controller.text.trim();
    if (BaseUtil.isEmpty(phoneNumber)) {
      ToastX.show('请输入手机号');
      return;
    }
    if (phoneNumber.length != 11) {
      ToastX.show('请输入正确的手机号');
      return;
    }
    setState(() {
      isLoading = true;
    });
    bool result = await AuthService.getSendSmsCode(phoneNumber);
    setState(() {
      isLoading = false;
    });
    if (result) {
      go(
        Routes.inputCode,
        arguments: InputCodeArgument(phone: phoneNumber),
      );
    }
  }
}
