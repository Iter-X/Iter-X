import 'package:client/app/events/events.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/toast.dart';
import 'package:dio/dio.dart';

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
    final responseCode = response.statusCode;
    final responseResult = response.data;
    HttpResultBean resultData;
    try {
      if (responseCode == 200) {
        resultData = HttpResultBean(code: responseCode, data: responseResult);
      } else {
        if (responseCode == 401) {
          eventBus.fire(EventUnauthorized(code: 401));
          ToastX.show('登录信息失效');
        } else if (responseCode == 500) {
          ToastX.show('服务器异常');
        }
        resultData = HttpResultBean.fromJson(responseResult);
      }
    } catch (exception) {
      resultData = HttpResultBean(code: response.statusCode, msg: "解析错误");
    }
    return resultData;
  }
}
