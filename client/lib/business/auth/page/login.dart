import 'package:client/app/constants.dart';
import 'package:client/app/routes.dart';
import 'package:client/common/material/image.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/toast.dart';
import 'package:client/common/widgets/base_button.dart';
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
    color: AppColor.primary,
    fontSize: 14.sp,
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
          Container(
            margin: EdgeInsets.symmetric(horizontal: 72.w),
            child: BaseButton(
              iconName: 'ic_wx.png',
              text: '微信登录',
              textColor: AppColor.secondary,
              backgroundColor: AppColor.primary,
              onTap: () => loginType(0),
            ),
          ),
          Gap(15.h),
          Container(
            margin: EdgeInsets.symmetric(horizontal: 72.w),
            child: BaseButton(
              iconName: 'ic_phone.png',
              text: '手机登录',
              hasBorder: true,
              onTap: () => loginType(1),
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
                Gap(5.w),
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
                top: 35.h,
                bottom: 35.h,
              ),
              decoration: BoxDecoration(
                borderRadius: BorderRadius.circular(AppConfig.boxRadius),
                color: AppColor.secondary,
              ),
              child: Column(
                children: [
                  Text(
                    '阅读以下协议并同意',
                    style: TextStyle(
                      color: AppColor.primary,
                      fontSize: 18.sp,
                      fontWeight: AppFontWeight.semiBold,
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
                  Container(
                    margin: EdgeInsets.only(
                      top: 20.h,
                      left: 32.w,
                      right: 32.w,
                    ),
                    child: BaseButton(
                      text: '同意',
                      height: 42.h,
                      textColor: Colors.white,
                      backgroundColor: AppColor.primary,
                      onTap: () {
                        setState(() {
                          isSelect = !isSelect;
                        });
                        Navigator.pop(context);
                        wxOrPhone(type);
                      },
                    ),
                  ),
                  Container(
                    margin: EdgeInsets.only(
                      top: 10.h,
                      left: 32.w,
                      right: 32.w,
                    ),
                    child: BaseButton(
                      text: '取消',
                      height: 42.h,
                      hasBorder: true,
                      onTap: () => Navigator.pop(context),
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
      // TODO: wechat login
      ToastX.showTODO();
    } else {
      go(Routes.phoneLogin);
    }
  }
}
