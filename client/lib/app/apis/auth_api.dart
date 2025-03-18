class AuthApi {
  // 获取验证码
  static const String sendSmsCode = '/api/v1/auth/send-sms-code';
  // 验证码登录
  static const String verifyLogin = '/api/v1/auth/verify-sms-code';
  // 一键登录
  static const String oneClickLogin = '/api/v1/auth/one-click-login';
}