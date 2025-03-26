class AuthApi {
  // 获取验证码
  static const String sendSmsCode = '/api/v1/auth/send-sms-code';
  // 验证码登录
  static const String verifyLogin = '/api/v1/auth/verify-sms-code';
  // 一键登录
  static const String oneClickLogin = '/api/v1/auth/one-click-login';
  // 获取用户信息
  static const String getUserInfo = '/api/v1/user/info';
  // 刷新 token
  static const String refreshToken = '/api/v1/auth/refresh-token';
}