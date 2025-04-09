class Collaborator {
  final String id;
  final String username;
  final String nickname;
  final String avatarUrl;

  static const String defaultAvatar = 'assets/images/placeholder.png';

  Collaborator({
    required this.id,
    required this.username,
    required this.nickname,
    required this.avatarUrl,
  });

  factory Collaborator.fromJson(Map<String, dynamic> json) {
    return Collaborator(
      id: json['id'] as String,
      username: json['username'] as String,
      nickname: (json['nickname'] as String?) ?? '',
      avatarUrl: (json['avatarUrl'] as String?) ?? defaultAvatar,
    );
  }
}
