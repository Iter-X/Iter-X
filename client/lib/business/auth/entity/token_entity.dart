class TokenEntity {
  final String token;
  final String? refreshToken;
  final int? expiresIn;
  final int? expiresAt; // 过期时间戳（毫秒）

  TokenEntity({
    required this.token,
    this.refreshToken,
    this.expiresIn,
    this.expiresAt,
  });

  factory TokenEntity.fromJson(Map<String, dynamic> json) {
    return TokenEntity(
      token: json['token'] ?? '',
      refreshToken: json['refreshToken'],
      expiresIn: json['expiresIn']?.toInt(),
      expiresAt: json['expiresAt']?.toInt(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'token': token,
      'refreshToken': refreshToken,
      'expiresIn': expiresIn,
      'expiresAt': expiresAt,
    };
  }
} 