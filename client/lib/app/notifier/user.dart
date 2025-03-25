import 'dart:convert';

import 'package:client/common/utils/logger.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

import '../../business/auth/entity/user_info_entity.dart';
import '../../common/material/state.dart';
import '../events/events.dart';

const String _kToken = 'token';

class UserNotifier with ChangeNotifier, DiagnosticableTreeMixin {
  final FlutterSecureStorage _storage = FlutterSecureStorage();

  UserInfoEntity? _user;

  UserInfoEntity? get user => _user;

  bool get isLoggedIn => _user != null; // need to confirm if the token is expired

  Future<void> loadUserInfo() async {
    String? token = await _storage.read(key: _kToken);
    BaseLogger.v('load token from storage: $token');
    if (token != null) {
      _user = UserInfoEntity(
        token: token,
      );
      notifyListeners();
    }
  }

  Future<void> login(UserInfoEntity user) async {
    if (_user == user) {
      return;
    }

    _user = user;

    await _storage.write(key: _kToken, value: user.token);

    notifyListeners();
    eventBus.fire(EventUserLoginStatusChange(true));
  }

  Future<void> logout() async {
    _user = null;

    await _storage.delete(key: _kToken);

    notifyListeners();
    eventBus.fire(EventUserLoginStatusChange(false));
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(StringProperty('user', json.encode(_user?.toJson())));
  }
}
