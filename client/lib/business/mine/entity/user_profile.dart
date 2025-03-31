import 'package:client/business/auth/entity/user_info_entity.dart';

class UserProfile {
  final UserInfoEntity userInfo;
  final int exploredCountries;
  final int exploredCities;
  final int exploredStates;
  final int completedBucketListItems;
  final int totalBucketListItems;
  final List<TripPreview> recentTrips;

  UserProfile({
    required this.userInfo,
    required this.exploredCountries,
    required this.exploredCities,
    required this.exploredStates,
    required this.completedBucketListItems,
    required this.totalBucketListItems,
    required this.recentTrips,
  });

  UserProfile.empty()
      : userInfo = UserInfoEntity(
          id: '',
          username: '',
          nickname: '',
          email: '',
          phoneNumber: '',
          avatarUrl: '',
          createdAt: '',
          updatedAt: '',
        ),
        exploredCountries = 0,
        exploredCities = 0,
        exploredStates = 0,
        completedBucketListItems = 0,
        totalBucketListItems = 0,
        recentTrips = [];

  factory UserProfile.fromJson(Map<String, dynamic> json) {
    return UserProfile(
      userInfo:
          UserInfoEntity.fromJson(json['userInfo'] as Map<String, dynamic>),
      exploredCountries: json['exploredCountries'] as int,
      exploredCities: json['exploredCities'] as int,
      exploredStates: json['exploredStates'] as int,
      completedBucketListItems: json['completedBucketListItems'] as int,
      totalBucketListItems: json['totalBucketListItems'] as int,
      recentTrips: (json['recentTrips'] as List<dynamic>)
          .map((e) => TripPreview.fromJson(e as Map<String, dynamic>))
          .toList(),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'userInfo': userInfo.toJson(),
      'exploredCountries': exploredCountries,
      'exploredCities': exploredCities,
      'exploredStates': exploredStates,
      'completedBucketListItems': completedBucketListItems,
      'totalBucketListItems': totalBucketListItems,
      'recentTrips': recentTrips.map((e) => e.toJson()).toList(),
    };
  }
}

class TripPreview {
  final String title;
  final String imageUrl;

  TripPreview({
    required this.title,
    required this.imageUrl,
  });

  TripPreview.empty()
      : title = '',
        imageUrl = '';

  factory TripPreview.fromJson(Map<String, dynamic> json) {
    return TripPreview(
      title: json['title'] as String,
      imageUrl: json['imageUrl'] as String,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'title': title,
      'imageUrl': imageUrl,
    };
  }
}
