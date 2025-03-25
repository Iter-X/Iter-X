import 'package:client/common/dio/http_result_bean.dart';

import '../../../app/apis/auth_api.dart';
import '../../../common/dio/http.dart';
import '../entity/token_entity.dart';
import '../entity/user_info_entity.dart';

class AuthService {
  // 获取验证码
  static Future<bool> getSendSmsCode(String phoneNumber) async {
    HttpResultBean result = await Http.instance.post(
      AuthApi.sendSmsCode,
      data: {
        'phoneNumber': phoneNumber,
      },
    );
    return result.isSuccess();
  }

  // 验证码登录
  static Future<TokenEntity?> verifyLogin(
    String phoneNumber,
    String verifyCode,
  ) async {
    HttpResultBean result = await Http.instance.post(
      AuthApi.verifyLogin,
      data: {
        'phoneNumber': phoneNumber,
        'verifyCode': verifyCode,
      },
    );
    if (result.isSuccess()) {
      return TokenEntity.fromJson(result.data);
    }
    return null;
  }

  // 一键登录
  static Future<TokenEntity?> oneClickLogin(String token) async {
    HttpResultBean result = await Http.instance.post(
      AuthApi.oneClickLogin,
      data: {
        'token': token,
      },
    );
    if (result.isSuccess()) {
      return TokenEntity.fromJson(result.data);
    }
    return null;
  }

  // 获取用户信息
  static Future<UserInfoEntity?> getUserInfo() async {
    HttpResultBean result = await Http.instance.get(
      '/api/v1/user/info',
    );
    if (result.isSuccess()) {
      return UserInfoEntity.fromJson(result.data);
    }
    return null;
  }

  // 刷新 token
  static Future<TokenEntity?> refreshToken(String refreshToken) async {
    HttpResultBean result = await Http.instance.post(
      '/api/v1/auth/refresh-token',
      data: {
        'refreshToken': refreshToken,
      },
    );
    if (result.isSuccess()) {
      return TokenEntity.fromJson(result.data);
    }
    return null;
  }
}
