import 'dart:collection';
import 'dart:convert';
import 'dart:io';

import 'package:client/app/events/events.dart';
import 'package:client/app/notifier/user.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/api_util.dart';
import 'package:client/common/utils/logger.dart';
import 'package:client/common/utils/util.dart';
import 'package:client/main.dart';
import 'package:dio/dio.dart';
import 'package:flutter/widgets.dart';
import 'package:path/path.dart' as path;
import 'package:path_provider/path_provider.dart';
import 'package:permission_handler/permission_handler.dart';
import 'package:provider/provider.dart';

import 'app_network_config.dart';
import 'http_result_bean.dart';

class HttpMethod {
  final String _name;

  const HttpMethod._(this._name);

  @override
  String toString() {
    return _name;
  }

  static const HttpMethod get = HttpMethod._('GET');
  static const HttpMethod post = HttpMethod._('POST');
}

class HttpUtil {
  static Map<String, dynamic> addIfNotEmpty(
    Map<String, dynamic> map,
    String key,
    dynamic value,
  ) {
    if (value == null) {
      return map;
    }
    if (value is String) {
      if (value.isNotEmpty) {
        map[key] = value;
      }
    } else {
      map[key] = value;
    }
    return map;
  }
}

class Http {
  static late AppNetworkConfig _sInstanceConfig = AppNetworkConfig();

  /*
   * 必须调此方法做初始化单例
   * 业务层自定义网络配置
   */
  static void initConfig(AppNetworkConfig config) {
    _sInstanceConfig = config;
  }

  static final Http _instance = Http._internal(_sInstanceConfig);

  static Http get instance => _instance;

  factory Http() => _instance;

  Dio dio = Dio();

  AppNetworkConfig get config => _sInstanceConfig;

  // 传入AppNetworkConfig的构造方法
  Http._internal(AppNetworkConfig config) {
    dio.options = BaseOptions(
      connectTimeout: Duration(milliseconds: config.connectTimeout),
      receiveTimeout: Duration(milliseconds: config.receiveTimeout),
      sendTimeout: Duration(milliseconds: config.sendTimeout),
      baseUrl: ApiUtil.apiModel.baseUrl ?? '',
      validateStatus: (int? status) => status != null && status > 0,
      headers: config.headers ?? {},
    );
    if (BaseLogger.enable) {
      dio.interceptors.add(LogInterceptor());
    }
    dio.interceptors.add(
      QueuedInterceptorsWrapper(
        onRequest: (options, handler) async {
          String apiPath =
              options.path.substring(0, options.path.lastIndexOf('/'));

          // 不需要 token 的接口直接放行
          if (apiPath.contains('/auth/') || options.path.contains('/auth/')) {
            return handler.next(options);
          }

          // 获取 UserNotifier 实例
          final BuildContext? context = navigatorKey.currentContext;
          if (context == null) {
            return handler.reject(
              DioException(
                requestOptions: options,
                error: 'No context available',
              ),
            );
          }

          final userNotifier = context.read<UserNotifier>();

          // 检查是否有有效的 token
          if (!await userNotifier.ensureValidToken()) {
            return handler.reject(
              DioException(
                requestOptions: options,
                error: 'Token expired or invalid',
              ),
            );
          }

          // 添加 token 到请求头
          if (userNotifier.token != null) {
            options.headers['Authorization'] =
                'Bearer ${userNotifier.token!.token}';
          }

          options.headers['X-Lang'] = 'zh-CN';

          BaseLogger.i('request headers: ${json.encode(options.headers)}');
          try {
            BaseLogger.i('request data: ${json.encode(options.data)}');
          } catch (exception) {
            BaseLogger.i('request data: ${options.data}');
          }
          return handler.next(options);
        },
        onError: (error, handler) async {
          if (error.response?.statusCode == 401) {
            final BuildContext? context = navigatorKey.currentContext;
            if (context != null) {
              final userNotifier = context.read<UserNotifier>();
              // 触发未授权事件
              eventBus.fire(EventUnauthorized());
            }
          }
          return handler.next(error);
        },
      ),
    );
    var interceptor = config.interceptor;
    if (BaseUtil.isNotEmpty(interceptor)) {
      dio.interceptors.addAll(interceptor);
    }
  }

  // 域名切换重试
  void init({required String baseUrl}) {
    //源码会报错会报错  https://github.com/flutterchina/dio/issues/1133
    // 'You cannot set both contentType param and a content-type header',
    //会设置两次contentType  参考 option   651行  所以必须要移除
    dio.options = dio.options.copyWith(
        headers: Map.from(dio.options.headers)
          ..remove(Headers.contentTypeHeader),
        baseUrl: baseUrl);
  }

  // 响应数据統一转换
  Future<HttpResultBean> _handleResponseData(
      Response response, bool isShowLoading) async {
    hideLoading(isShowLoading);
    //回到APP业务层去处理
    HttpResultBean result =
        await Http.instance.config.handleResponseData(response);
    return result;
  }

  // 公共请求参数
  Future<Map<String, dynamic>> getBasicParam() async {
    Map<String, dynamic> basicParam = {};
    // var uuid = Uuid();
    // basicParam["uuid"] = uuid.v4();
    return basicParam;
  }

  // restful get 操作
  Future get(
    String path, {
    Map<String, dynamic>? params,
    Options? options,
    CancelToken? cancelToken,
    bool isShowLoading = false,
    ProgressCallback? onReceiveProgress,
  }) async {
    params ??= HashMap<String, dynamic>();
    showLoading(isShowLoading);
    try {
      Response<Map<String, dynamic>> response =
          await dio.get<Map<String, dynamic>>(
        path,
        queryParameters: params,
        options: options,
        cancelToken: cancelToken,
        onReceiveProgress: onReceiveProgress,
      );
      BaseLogger.i('$path / response ${json.encode(response.data)}');
      return _handleResponseData(response, isShowLoading);
    } catch (exception) {
      BaseLogger.e("----catch--请求错误-----${exception}");
      hideLoading(isShowLoading);
      return HttpResultBean(code: 404, msg: "网络超时");
    }
  }

  // restful post 操作
  Future post(
    String path, {
    Map<String, dynamic>? data,
    dataList,
    bool isShowLoading = false,
    Map<String, dynamic>? queryParameters,
    Options? options,
    CancelToken? cancelToken,
    ProgressCallback? onSendProgress,
    ProgressCallback? onReceiveProgress,
  }) async {
    // 传参统一处理, 加上基本参数
    Map<String, dynamic> basicParam = await getBasicParam();
    if (data != null) {
      basicParam.addAll(data);
    }
    showLoading(isShowLoading);
    try {
      Response<Map<String, dynamic>> response =
          await dio.post<Map<String, dynamic>>(
        path,
        data: dataList ?? basicParam,
        queryParameters: queryParameters,
        options: options,
        cancelToken: cancelToken,
        onSendProgress: onSendProgress,
        onReceiveProgress: onReceiveProgress,
      );
      BaseLogger.i('response ${json.encode(response.data)}');
      return _handleResponseData(response, isShowLoading);
    } catch (exception) {
      BaseLogger.e("--catch--请求错误-----$exception");
      hideLoading(isShowLoading);
      return HttpResultBean(code: 404, msg: "网络超时");
    }
  }

  // restful post form 表单提交操作
  Future postForm(
    String path, {
    bool isShowLoading = false,
    Map<String, dynamic> data = const {},
    Map<String, dynamic>? queryParameters,
    Options? options,
    CancelToken? cancelToken,
    ProgressCallback? onSendProgress,
    ProgressCallback? onReceiveProgress,
  }) async {
    showLoading(isShowLoading);
    Response<Map<String, dynamic>> response =
        await dio.post<Map<String, dynamic>>(
      path,
      data: FormData.fromMap(data),
      queryParameters: queryParameters,
      options: options,
      cancelToken: cancelToken,
      onSendProgress: onSendProgress,
      onReceiveProgress: onReceiveProgress,
    );
    BaseLogger.i('response ${json.encode(response.data)}');
    return _handleResponseData(response, isShowLoading);
  }

  // 文件提交
  Future uploadFile(
    String path, {
    required Map<String, dynamic> data,
    ProgressCallback? onSendProgress,
    CancelToken? cancelToken,
  }) async {
    Response response = await dio.post(
      path,
      data: FormData.fromMap(data),
      options: Options(
        contentType: "multipart/form-data",
        sendTimeout: const Duration(milliseconds: 120000),
        receiveTimeout: const Duration(milliseconds: 120000),
      ),
      onSendProgress: onSendProgress,
      cancelToken: cancelToken,
    );
    BaseLogger.i('response ${json.encode(response.data)}');
    return _handleResponseData(response, false);
  }

  // 文件下载
  Future downFile(String path) async {
    Response<String> response = await dio.post<String>(path);
    BaseLogger.i('response ${json.encode(response.data)}');
    return _handleResponseData(response, false);
  }

  // restful put 操作
  Future put(
    String path, {
    data,
    Map<String, dynamic>? queryParameters,
    Options? options,
    CancelToken? cancelToken,
    ProgressCallback? onSendProgress,
    ProgressCallback? onReceiveProgress,
  }) async {
    try {
      Response<Map<String, dynamic>> response =
          await dio.put<Map<String, dynamic>>(
        path,
        data: data,
        queryParameters: queryParameters,
        options: options,
        cancelToken: cancelToken,
        onSendProgress: onSendProgress,
        onReceiveProgress: onReceiveProgress,
      );
      BaseLogger.i('response ${json.encode(response.data)}');
      return _handleResponseData(response, false);
    } catch (exception) {
      BaseLogger.e("----catch--请求错误-----${exception}");
      return HttpResultBean(code: 404, msg: "网络超时");
    }
  }

  // restful patch 操作
  Future patch(
    String path, {
    data,
    Map<String, dynamic>? queryParameters,
    Options? options,
    CancelToken? cancelToken,
    ProgressCallback? onSendProgress,
    ProgressCallback? onReceiveProgress,
  }) async {
    Response<Map<String, dynamic>> response =
        await dio.patch<Map<String, dynamic>>(
      path,
      data: data,
      queryParameters: queryParameters,
      options: options,
      cancelToken: cancelToken,
      onSendProgress: onSendProgress,
      onReceiveProgress: onReceiveProgress,
    );
    BaseLogger.i('response ${json.encode(response.data)}');
    return response.data;
  }

  // restful delete 操作
  Future delete(
    String path, {
    data,
    Map<String, dynamic>? queryParameters,
    Options? options,
    CancelToken? cancelToken,
    ProgressCallback? onSendProgress,
    ProgressCallback? onReceiveProgress,
  }) async {
    try {
      Response<Map<String, dynamic>> response =
          await dio.delete<Map<String, dynamic>>(
        path,
        data: data,
        queryParameters: queryParameters,
        options: options,
        cancelToken: cancelToken,
      );
      BaseLogger.i('response ${json.encode(response.data)}');
      return _handleResponseData(response, false);
    } catch (exception) {
      BaseLogger.e("----catch--请求错误-----${exception}");
      return HttpResultBean(code: 404, msg: "网络超时");
    }
  }

  static void showLoading(bool isShow) {
    // if (isShow) {
    //   FlutterLoading.show(dismissOnTap: false);
    // }
  }

  static void hideLoading(bool isShow) {
    // if (isShow) {
    //   FlutterLoading.dismiss();
    // }
  }

  Future<String?> download(
    String url, {
    ProgressCallback? onReceiveProgress,
    CancelToken? cancelToken,
  }) async {
    try {
      Directory tempDir = await getTemporaryDirectory();
      String savePath = path.join(tempDir.path, path.basename(url));

      bool permission = await requestStoragePermission();
      if (!permission) {
        return null;
      }

      await dio.download(
        url,
        savePath,
        cancelToken: cancelToken,
        onReceiveProgress: onReceiveProgress,
      );

      return savePath;
    } catch (exception) {
      BaseLogger.e("----catch--请求错误-----${exception}");
      return null;
    }
  }

  /// 申请定位权限
  /// 授予定位权限返回true， 否则返回false
  static Future<bool> requestStoragePermission() async {
    //获取当前的权限
    var status = await Permission.storage.status;
    if (status == PermissionStatus.granted) {
      //已经授权
      return true;
    } else {
      //未授权则发起一次申请
      status = await Permission.storage.request();
      if (status == PermissionStatus.granted) {
        return true;
      } else {
        return false;
      }
    }
  }
}
