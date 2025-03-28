import 'package:client/app/routes.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/color.dart';
import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';


class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends BaseState<LoginPage> {
  TextStyle agreementTextStyle = TextStyle(
    color: BaseColor.c_1D1F1E,
    fontSize: 16.sp,
  );

  bool isSelect = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        mainAxisAlignment: MainAxisAlignment.end,
        children: [
          BaseImage.asset(
            name: 'ic_logo.png',
            width: 245.w,
          ),
          Gap(200.h),
          GestureDetector(
            onTap: () => loginType(0),
            child: Container(
              width: 285.w,
              height: 52.h,
              decoration: BoxDecoration(
                color: BaseColor.c_1D1F1E,
                borderRadius: BorderRadius.circular(24.r),
              ),
              padding: EdgeInsets.only(left: 37.w),
              child: Row(
                children: [
                  BaseImage.asset(
                    name: 'ic_wx.png',
                    size: 30.w,
                  ),
                  Gap(49.w),
                  Text(
                    '微信登录',
                    style: TextStyle(
                      color: BaseColor.c_F2F2F2,
                      fontSize: 20.sp,
                    ),
                  )
                ],
              ),
            ),
          ),
          GestureDetector(
            onTap: () => loginType(1),
            child: Container(
              width: 285.w,
              height: 52.h,
              margin: EdgeInsets.only(top: 15.h),
              decoration: BoxDecoration(
                color: Colors.transparent,
                borderRadius: BorderRadius.circular(24.r),
                border: Border.all(
                  width: 1.w,
                  color: BaseColor.c_1D1F1E,
                ),
              ),
              padding: EdgeInsets.only(left: 37.w),
              child: Row(
                children: [
                  BaseImage.asset(
                    name: 'ic_phone.png',
                    size: 30.w,
                  ),
                  Gap(49.w),
                  Text(
                    '手机登录',
                    style: TextStyle(
                      color: BaseColor.c_1D1F1E,
                      fontSize: 20.sp,
                    ),
                  )
                ],
              ),
            ),
          ),
          Container(
            margin: EdgeInsets.only(
              left: 44.w,
              right: 44.w,
              top: 107.h,
              bottom: 73.h,
            ),
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                GestureDetector(
                  onTap: () {
                    setState(() {
                      isSelect = !isSelect;
                    });
                  },
                  child: Container(
                    margin: EdgeInsets.only(top: 2.h),
                    child: BaseImage.asset(
                      name: isSelect ? 'ic_select.png' : 'ic_unselect.png',
                      size: 20.w,
                    ),
                  ),
                ),
                Gap(12.w),
                Expanded(
                  child: RichText(
                    text: TextSpan(
                      children: [
                        TextSpan(
                          text: '我已阅读并同意',
                          style: agreementTextStyle,
                        ),
                        TextSpan(
                          text: '《用户协议》',
                          style: agreementTextStyle,
                          recognizer: TapGestureRecognizer()..onTap = () {},
                        ),
                        TextSpan(
                          text: '、',
                          style: agreementTextStyle,
                        ),
                        TextSpan(
                          text: '《隐私政策》',
                          style: agreementTextStyle,
                          recognizer: TapGestureRecognizer()..onTap = () {},
                        ),
                        TextSpan(
                          text: '和',
                          style: agreementTextStyle,
                        ),
                        TextSpan(
                          text: '《儿童/青少年个人信息保护规则》',
                          style: agreementTextStyle,
                          recognizer: TapGestureRecognizer()..onTap = () {},
                        ),
                      ],
                    ),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }

  // 微信登录
  void loginType(int type) {
    if (!isSelect) {
      showDialog(
        context: context,
        builder: (builder) {
          return UnconstrainedBox(
            constrainedAxis: Axis.horizontal,
            child: Container(
              width: double.infinity,
              margin: EdgeInsets.symmetric(horizontal: 60.w),
              padding: EdgeInsets.only(
                top: 29.h,
                bottom: 26.h,
              ),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(12.w),
                color: BaseColor.c_F2F2F2,
              ),
              child: Column(
                children: [
                  Text(
                    '阅读以下协议并同意',
                    style: TextStyle(
                      color: BaseColor.c_1D1F1E,
                      fontSize: 20.sp,
                      fontWeight: FontWeight.w600,
                    ),
                  ),
                  Container(
                    margin: EdgeInsets.only(top: 20.h),
                    child: RichText(
                      text: TextSpan(
                        children: [
                          TextSpan(
                            text: '《用户协议》',
                            style: agreementTextStyle,
                            recognizer: TapGestureRecognizer()..onTap = () {},
                          ),
                          TextSpan(
                            text: '《隐私政策》',
                            style: agreementTextStyle,
                            recognizer: TapGestureRecognizer()..onTap = () {},
                          ),
                        ],
                      ),
                    ),
                  ),
                  GestureDetector(
                    onTap: () {},
                    child: Text(
                      '《儿童/青少年个人信息保护规则》',
                      style: agreementTextStyle,
                    ),
                  ),
                  GestureDetector(
                    onTap: () {
                      setState(() {
                        isSelect = !isSelect;
                      });
                      Navigator.pop(context);
                      wxOrPhone(type);
                    },
                    child: Container(
                      width: double.infinity,
                      height: 42.h,
                      margin: EdgeInsets.only(
                        left: 31.w,
                        right: 31.w,
                        top: 20.h,
                      ),
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(24.w),
                        color: BaseColor.c_1D1F1E,
                      ),
                      alignment: Alignment.center,
                      child: Text(
                        '同意',
                        style: TextStyle(
                          color: Colors.white,
                          fontSize: 20.sp,
                        ),
                      ),
                    ),
                  ),
                  GestureDetector(
                    onTap: () => Navigator.pop(context),
                    child: Container(
                      width: double.infinity,
                      height: 42.h,
                      margin: EdgeInsets.only(
                        left: 31.w,
                        right: 31.w,
                        top: 11.h,
                      ),
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(24.w),
                        color: Colors.transparent,
                        border: Border.all(
                          width: 1.w,
                          color: BaseColor.c_1D1F1E,
                        ),
                      ),
                      alignment: Alignment.center,
                      child: Text(
                        '取消',
                        style: TextStyle(
                          color: BaseColor.c_1D1F1E,
                          fontSize: 20.sp,
                        ),
                      ),
                    ),
                  ),
                ],
              ),
            ),
          );
        },
      );
    } else {
      wxOrPhone(type);
    }
  }

  // 微信登录or手机号登录
  void wxOrPhone(int type) {
    if (type == 0) {

    } else {
      go(Routes.phoneLogin);
    }
  }
}
