import 'dart:convert';

import 'package:client/app/constants.dart';
import 'package:client/app/routes.dart';
import 'package:client/business/auth/page/input_code.dart';
import 'package:client/business/auth/service/auth_service.dart';
import 'package:client/common/material/loading.dart';
import 'package:client/common/material/text_field.dart';
import 'package:client/common/utils/color.dart';
import 'package:client/common/utils/toast.dart';
import 'package:client/common/utils/util.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:gap/gap.dart';
import 'package:ali_auth/ali_auth.dart';

import '../../../common/dio/http_result_bean.dart';
import '../../../common/material/state.dart';
import '../../../common/utils/shared_preference_util.dart';

class PhoneLoginPage extends StatefulWidget {
  const PhoneLoginPage({super.key});

  @override
  State<PhoneLoginPage> createState() => _PhoneLoginPageState();
}

class _PhoneLoginPageState extends BaseState<PhoneLoginPage> {
  late TextEditingController _controller;
  bool isLoading = false;

  @override
  void initState() {
    _controller = TextEditingController();
    initAliAuth();
    super.initState();
  }

  @override
  void dispose() {
    _controller.dispose();
    AliAuth.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      resizeToAvoidBottomInset: false,
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
                child: isLoading
                    ? const LoadingWidget()
                    : Text(
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

  Future<void> codeLogin() async {
    if (isLoading) {
      return;
    }
    String phoneNumber = _controller.text.trim();
    if (BaseUtil.isEmpty(phoneNumber)) {
      Toast.show('请输入手机号');
      return;
    }
    if (phoneNumber.length != 11) {
      Toast.show('请输入正确的手机号');
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

  Future<void> initAliAuth() async {
    await AliAuth.loginListen(
      onEvent: (onEvent) {
        print("----------------> $onEvent <----------------");
        Map<String, dynamic> response = jsonDecode(jsonEncode(onEvent));
        if (response['code'] == '600000') {
          // 登录成功
          oneClickLogin(response['data']);
        }
      },
      isOnlyOne: false,
    );
    await AliAuth.initSdk(getFullPortConfig());
  }

  /// 全屏正常图片背景
  AliAuthModel getFullPortConfig({bool isDelay = false}) {
    return AliAuthModel(
      Constants.aliAndroidSdk,
      Constants.aliIosSdk,
      isDebug: true,
      autoQuitPage: true,
      isDelay: isDelay,
      pageType: PageType.fullPort,
      statusBarColor: "#f2f2f2",
      lightColor: true,
      isStatusBarHidden: false,
      statusBarUIFlag: UIFAG.systemUiFalgFullscreen,
      isHiddenCustom: false,
      customReturnBtn: CustomView(
        59,
        0,
        0,
        26,
        28,
        28,
        'assets/images/return_btn.png',
        ScaleType.fitCenter,
      ),
      navReturnImgWidth: 30,
      navReturnImgHeight: 30,
      navReturnHidden: true,
      navReturnScaleType: ScaleType.center,
      navHidden: true,
      navTextSize: 18,
      navText: '',
      navColor: "#f2f2f2",
      backgroundColor: '#f2f2f2',
      // logoOffsetY: 392,
      // logoImgPath: "assets/logo.png",
      // logoHidden: false,
      logBtnHeight: 52,
      logBtnOffsetX: 0,
      logBtnOffsetY: 394,
      // logBtnMarginLeftAndRight: 73,
      logBtnLayoutGravity: Gravity.centerHorizntal,
      logBtnWidth: 285,
      logBtnText: "本机一键登录",
      logBtnTextSize: 18,
      logBtnTextColor: "#F2F2F2",
      logBtnBackgroundPath: "assets/images/login_btn_normal.png",
      // protocolOneName: "《通达理》",
      // protocolOneURL: "https://tunderly.com",
      // protocolTwoName: "《思预云》",
      // protocolTwoURL: "https://jokui.com",
      // protocolThreeName: "《思预云APP》",
      // protocolThreeURL:
      // "https://a.app.qq.com/o/simple.jsp?pkgname=com.civiccloud.master&fromcase=40002",
      // protocolCustomColor: "#026ED2",
      // protocolColor: "#bfbfbf",
      // protocolLayoutGravity: Gravity.centerHorizntal,
      numberColor: "#1D1F1E",
      numberSize: 28,
      numFieldOffsetY: 292,
      numberFieldOffsetX: 0,
      numberLayoutGravity: Gravity.centerHorizntal,
      privacyOffsetX: -1,
      privacyOffsetY: -1,
      privacyOffsetY_B: 28,
      checkBoxWidth: 18,
      checkBoxHeight: 18,
      checkboxHidden: false,
      privacyState: false,
      switchAccTextSize: 16,
      switchAccText: "切换到其他方式",
      switchOffsetY_B: -1,
      switchAccHidden: false,
      switchAccTextColor: "#FDFDFD",
      sloganTextSize: 16,
      sloganHidden: false,
      uncheckedImgPath: "assets/btn_unchecked.png",
      checkedImgPath: "assets/btn_checked.png",
      protocolGravity: Gravity.centerHorizntal,
      privacyTextSize: 12,
      privacyMargin: 28,
      privacyBefore: "已阅读并同意",
      privacyEnd: "",
      vendorPrivacyPrefix: "《",
      vendorPrivacySuffix: "》",
      dialogWidth: -1,
      dialogHeight: -1,
      dialogBottom: false,
      dialogOffsetX: 0,
      dialogOffsetY: 0,
      // webViewStatusBarColor: "",
      // webNavColor: "#026ED2",
      // webNavTextColor: "#ffffff",
      // webNavTextSize: 20,
      // webNavReturnImgPath: "assets/return_btn.png",
      // webSupportedJavascript: true,
      authPageActIn: "in_activity",
      activityOut: "out_activity",
      authPageActOut: "in_activity",
      activityIn: "out_activity",
      screenOrientation: -1,
      logBtnToastHidden: false,
      dialogAlpha: 1.0,
      privacyOperatorIndex: 0,
      // customThirdView: customThirdView,
    );
  }

  // 一键登录
  void oneClickLogin(String token) async {
    var result = await AuthService.oneClickLogin(token);
    if (result != null) {
      await BaseSpUtil.setJSON(SpKeys.TOKEN, result.token);
      await BaseSpUtil.setJSON(SpKeys.USER_INFO, result);
      go(Routes.homeMain, clearStack: true);
    }
  }
}
