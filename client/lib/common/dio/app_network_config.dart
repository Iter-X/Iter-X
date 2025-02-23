import 'package:dio/dio.dart';

import '../utils/toast.dart';
import 'http_result_bean.dart';

/// 网络请求库配置
class AppNetworkConfig {
  late int connectTimeout;
  late int receiveTimeout;
  late int sendTimeout;
  late String? baseUrl;
  late bool retryInterceptorEnable;
  late Map<String, dynamic>? headers;
  late List<Interceptor> interceptor;

  AppNetworkConfig({
    this.connectTimeout = 20000,
    this.receiveTimeout = 20000,
    this.sendTimeout = 20000,
    this.baseUrl,
    this.headers,
    this.retryInterceptorEnable = true,
    this.interceptor = const [],
  });

  /*
   *  对响应的统计，处理全局的业务嘛，比如登录过期时候跳转登录页
   */
  Future<HttpResultBean> handleResponseData(Response response) async {
    final responseResult = response.data;
    HttpResultBean resultData;
    try {
      //dio code 0 仅仅标识网络请求 HTTP 请求码
      // if (responseResult["code"] == 417) {
      //   eventBus.fire(EventUnauthorized(code: 417));
      //   resultData = HttpResultBean(code: responseResult["code"], msg: "登录失效");
      // } else if (responseResult["code"] == 401) {
      //   eventBus.fire(EventUnauthorized(code: 401));
      //   resultData = HttpResultBean(code: responseResult["code"], msg: "登录信息失效");
      // } else if (response.statusCode != 200) {
      //   resultData = HttpResultBean(code: response.statusCode, msg: "网络请求错误");
      // } else {
        resultData = HttpResultBean.fromJson(responseResult);
      // }
    } catch (exception) {
      resultData = HttpResultBean(code: response.statusCode, msg: "解析错误");
    }
    if (resultData.isFail() && resultData.code != 10057) {
      Toast.show(resultData.msg ?? '网络请求错误');
    }
    return resultData;
  }
}
