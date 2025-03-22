// 登录状态事件
class EventUserLoginStatusChange {
  bool logined;

  EventUserLoginStatusChange(this.logined);
}

class EventUnauthorized {
  int? code;
  EventUnauthorized({this.code});
}
