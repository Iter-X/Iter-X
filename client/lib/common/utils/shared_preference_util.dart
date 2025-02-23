import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import 'util.dart';

class SpKeys {
  static String APIS = 'apis';
  static String TOKEN = 'token';
  static String USER_INFO = 'user_info';
}

class BaseSpUtil {
  /// 保存数据
  static Future<bool> save(String key, dynamic value) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    if (value is String) {
      return prefs.setString(key, value);
    } else if (value is int) {
      return prefs.setInt(key, value);
    } else if (value is double) {
      return prefs.setDouble(key, value);
    } else if (value is bool) {
      return prefs.setBool(key, value);
    } else if (value is List<String>) {
      return prefs.setStringList(key, value);
    } else {
      return remove(key);
    }
  }

  /// 删除指定key
  static Future<bool> remove(String key) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    return prefs.remove(key);
  }

  /// 清空
  static Future<bool> clear() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    return prefs.clear();
  }

  // 异步读取
  static Future<String> getString(String key,
      [String defaultValue = ""]) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    return sp.getString(key) ?? defaultValue;
  }

  // 异步读取
  static Future<String> getStringNull(String key,
      [String defaultValue = "null"]) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    return sp.getString(key) ?? defaultValue;
  }

  // 异步读取
  static Future<int> getInt(String key, [int defaultValue = 0]) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    return sp.getInt(key) ?? defaultValue;
  }

  // 异步读取
  static Future<double> getDouble(String key, [double defaultValue = 0]) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    return sp.getDouble(key) ?? defaultValue;
  }

  // 异步读取
  static Future<bool> getBool(String key, [bool defaultValue = false]) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    return sp.getBool(key) ?? defaultValue;
  }

  static Future<bool> setJSON(String key, dynamic jsonVal) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    String jsonString = jsonEncode(jsonVal);
    return sp.setString(key, jsonString);
  }

  static dynamic? getJSON(String key) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    String jsonString = sp.getString(key) ?? "";
    return BaseUtil.isEmpty(jsonString) ? null : jsonDecode(jsonString);
  }

  static Future<List<String>> getStringList(String key) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    List<String> jsonString = sp.getStringList(key) ?? [];
    return jsonString;
  }

  static Future<bool> isExist(String key) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    return sp.containsKey(key);
  }

  static Future<bool> containsKey(String key) async {
    SharedPreferences sp = await SharedPreferences.getInstance();
    return sp.containsKey(key);
  }
}
