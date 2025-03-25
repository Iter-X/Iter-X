// 登录状态事件
class EventUserLoginStatusChange {
  bool isLoggedIn;

  EventUserLoginStatusChange(this.isLoggedIn);
}

class EventUnauthorized {
  int? code;
  EventUnauthorized({this.code});
}
