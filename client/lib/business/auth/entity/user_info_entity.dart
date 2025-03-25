class UserInfoEntity {
  final String id;
  final String username;
  final String nickname;
  final String email;
  final String phoneNumber;
  final String avatarUrl;
  final String createdAt;
  final String updatedAt;

  UserInfoEntity({
    required this.id,
    required this.username,
    required this.nickname,
    required this.email,
    required this.phoneNumber,
    required this.avatarUrl,
    required this.createdAt,
    required this.updatedAt,
  });

  bool get hasValidAvatar => avatarUrl.isNotEmpty;

  factory UserInfoEntity.fromJson(Map<String, dynamic> json) {
    return UserInfoEntity(
      id: json['id'] as String,
      username: json['username'] as String,
      nickname: json['nickname'] as String,
      email: json['email'] as String,
      phoneNumber: json['phoneNumber'] as String,
      avatarUrl: json['avatarUrl'] as String,
      createdAt: json['createdAt'] as String,
      updatedAt: json['updatedAt'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'username': username,
      'nickname': nickname,
      'email': email,
      'phoneNumber': phoneNumber,
      'avatarUrl': avatarUrl,
      'createdAt': createdAt,
      'updatedAt': updatedAt,
    };
  }
}
