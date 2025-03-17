class UserInfoEntity {
  String? token;
  int? expiresIn;

  UserInfoEntity({
    this.token,
    this.expiresIn,
  });

  UserInfoEntity.fromJson(Map<String, dynamic> json) {
    token = json['token'];
    expiresIn = json['expiresIn'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = <String, dynamic>{};
    data['token'] = token;
    data['expiresIn'] = expiresIn;
    return data;
  }
}