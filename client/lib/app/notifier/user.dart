import 'dart:convert';

import 'package:client/app/events/events.dart';
import 'package:client/business/auth/entity/token_entity.dart';
import 'package:client/business/auth/entity/user_info_entity.dart';
import 'package:client/business/auth/service/auth_service.dart';
import 'package:client/common/material/state.dart';
import 'package:client/common/utils/logger.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';


const String _kToken = 'token';
const String _kUserInfo = 'user_info';

class UserNotifier with ChangeNotifier, DiagnosticableTreeMixin {
  final FlutterSecureStorage _storage = FlutterSecureStorage();

  UserInfoEntity? _user;
  TokenEntity? _token;

  UserInfoEntity? get user => _user;
  TokenEntity? get token => _token;

  // 检查 token 是否过期
  bool get isTokenExpired {
    if (_token == null || _token!.expiresAt == null) return true;
    
    final expirationTime = DateTime.fromMillisecondsSinceEpoch(_token!.expiresAt!);
    // 提前5分钟认为过期，以便有时间刷新
    return DateTime.now().isAfter(expirationTime.subtract(Duration(minutes: 5)));
  }

  Future<void> loadUserInfo() async {
    String? tokenStr = await _storage.read(key: _kToken);
    String? userInfoStr = await _storage.read(key: _kUserInfo);

    if (tokenStr != null) {
      try {
        Map<String, dynamic> tokenMap = json.decode(tokenStr);
        _token = TokenEntity.fromJson(tokenMap);
      } catch (e) {
        BaseLogger.e('Failed to parse token: $e');
      }
    }

    if (userInfoStr != null) {
      try {
        Map<String, dynamic> userInfoMap = json.decode(userInfoStr);
        _user = UserInfoEntity.fromJson(userInfoMap);
        notifyListeners();
      } catch (e) {
        BaseLogger.e('Failed to parse user info: $e');
      }
    }

    // 检查 token 是否过期，如果过期尝试刷新
    if (isTokenExpired) {
      await _handleTokenExpiration();
    }
  }

  // 处理 token 过期
  Future<bool> _handleTokenExpiration() async {
    if (_token?.refreshToken == null) {
      await logout();
      eventBus.fire(EventUnauthorized());
      return false;
    }

    try {
      final newToken = await AuthService.refreshToken(_token!.refreshToken!);
      if (newToken != null) {
        _token = newToken;
        await _storage.write(key: _kToken, value: json.encode(_token!.toJson()));
        notifyListeners();
        return true;
      }
    } catch (e) {
      BaseLogger.e('Failed to refresh token: $e');
    }

    await logout();
    eventBus.fire(EventUnauthorized());
    return false;
  }

  // 检查并确保 token 有效
  Future<bool> ensureValidToken() async {
    if (_token == null) return false;
    if (!isTokenExpired) return true;
    return _handleTokenExpiration();
  }

  Future<void> refreshUserInfo() async {
    BaseLogger.v('token: $token');
    BaseLogger.v('token valid: ${!isTokenExpired}');
    // 确保 token 有效
    if (!await ensureValidToken()) return;
    
    try {
      UserInfoEntity? updatedUser = await AuthService.getUserInfo();
      if (updatedUser != null) {
        _user = updatedUser;
        await _storage.write(key: _kUserInfo, value: json.encode(_user!.toJson()));
        notifyListeners();
      }
    } catch (e) {
      BaseLogger.e('Failed to refresh user info: $e');
    }
  }

  Future<void> login({
    required TokenEntity token,
  }) async {
    // 计算过期时间
    if (token.expiresIn != null) {
      _token = TokenEntity(
        token: token.token,
        refreshToken: token.refreshToken,
        expiresIn: token.expiresIn,
        expiresAt: DateTime.now().add(Duration(seconds: token.expiresIn!)).millisecondsSinceEpoch,
      );
    } else {
      _token = token;
    }

    BaseLogger.v('login token: ${_token!.toJson()}');
    
    await _storage.write(key: _kToken, value: json.encode(_token!.toJson()));

    notifyListeners();
    eventBus.fire(EventUserLoginStatusChange(true));
  }

  Future<void> updateUserInfo(UserInfoEntity user) async {
    _user = user;
    await _storage.write(key: _kUserInfo, value: json.encode(user.toJson()));
    notifyListeners();
  }

  Future<void> logout() async {
    _user = null;
    _token = null;

    await _storage.delete(key: _kToken);
    await _storage.delete(key: _kUserInfo);

    notifyListeners();
    eventBus.fire(EventUserLoginStatusChange(false));
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(StringProperty('user', json.encode(_user?.toJson())));
    properties.add(StringProperty('token', json.encode(_token?.toJson())));
  }
}
