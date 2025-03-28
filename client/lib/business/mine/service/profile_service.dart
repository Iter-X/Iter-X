import 'package:client/business/auth/entity/user_info_entity.dart';
import 'package:flutter/material.dart';

import '../entity/user_profile.dart';

class ProfileService extends ChangeNotifier {
  UserProfile? _profile;
  bool _isLoading = false;

  UserProfile? get profile => _profile;

  bool get isLoading => _isLoading;

  Future<void> fetchUserProfile() async {
    _isLoading = true;
    notifyListeners();

    try {
      // TODO: Replace with actual API call
      await Future.delayed(Duration(seconds: 1));
      _profile = UserProfile(
        userInfo: UserInfoEntity(
          id: '1',
          username: 'leo',
          nickname: 'Leo',
          email: 'ifuryst@gmail.com',
          phoneNumber: '1234567890',
          avatarUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/avatar.png',
          createdAt: DateTime.now().toIso8601String(),
          updatedAt: DateTime.now().toIso8601String(),
        ),
        exploredCountries: 27,
        exploredCities: 38,
        exploredStates: 3,
        completedBucketListItems: 105,
        totalBucketListItems: 2077,
        recentTrips: [
          TripPreview(
            title: '美西7天自驾游',
            imageUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/us.png',
          ),
          TripPreview(
            title: '成都2日游',
            imageUrl: 'https://nextstone.oss-cn-beijing.aliyuncs.com/cd.png',
          ),
        ],
      );
    } catch (e) {
      debugPrint('Error fetching user profile: $e');
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }
}
