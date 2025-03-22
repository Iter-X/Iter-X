class HttpResultBean {
  String? msg;
  int? code;
  dynamic data;

  HttpResultBean({
    this.msg,
    this.code,
    this.data,
  });

  HttpResultBean.fromJson(Map<String, dynamic> json) {
    msg = json['msg'];
    code = json['code'];
    data = json['data'] ?? "";
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> m = <String, dynamic>{};
    m['msg'] = msg;
    m['code'] = code;
    m['data'] = data;
    return m;
  }

  // 业务code 200 才是真正的成功
  bool isSuccess() {
    bool success = code == 200;
    return success;
  }

  bool isFail() {
    bool fail = code != 200;
    return fail;
  }
}
