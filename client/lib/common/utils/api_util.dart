import 'dart:convert';
import '../dio/api_model.dart';
import 'shared_preference_util.dart';

// API工具类
class ApiUtil {
  static const bool _isRelease = false;

  static ApiModel apiModel = _releaseApi;

  static final ApiModel _releaseApi = ApiModel(
    selected: _isRelease,
    title: '线上地址',
    baseUrl: 'https://api.iter-x.com',
  );

  static final List<ApiModel> _apis = [
    ApiModel(
      selected: !_isRelease,
      title: '测试地址',
      baseUrl: 'https://api.iter-x.com',
    ),
    _releaseApi,
  ];

  static Future<String> getBaseUrl() async {
    ApiModel api = await getSelectedApiModel();
    return '${api.baseUrl}';
  }

  static Future<ApiModel> getSelectedApiModel() async {
    List<ApiModel> apis = await getApis();
    for (var api in apis) {
      if (api.selected == true) {
        return api;
      }
    }
    return _releaseApi;
  }

  static Future<List<ApiModel>> getApis() async {
    String apis = await BaseSpUtil.getString(SpKeys.APIS);
    if (apis.isNotEmpty) {
      return (json.decode(apis) as List)
          .map((e) => ApiModel.fromJson(e))
          .toList();
    }
    return _apis;
  }
}
