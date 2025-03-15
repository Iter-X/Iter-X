import 'dart:convert';

import 'package:flutter/foundation.dart';

import '../../business/auth/entity/user_info_entity.dart';
import '../../common/material/state.dart';
import '../events/events.dart';

class UserNotifier with ChangeNotifier, DiagnosticableTreeMixin {
  UserNotifier(this._user);

  UserInfoEntity? _user;

  UserInfoEntity? get user => _user;

  bool logined() => _user != null;

  updateUserInfo(UserInfoEntity? user) {
    if (_user == user) {
      return;
    }
    if ((_user == null && user != null) || (_user != null && user == null)) {
      eventBus.fire(EventUserLoginStatusChange(user != null));
    }

    _user = user;
    notifyListeners();
  }

  updateUserInfoPart({
    String? token,
  }) {
    if (token != null) {
      _user?.token = token;
    }
    notifyListeners();
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(StringProperty('user', json.encode(_user?.toJson())));
  }
}
