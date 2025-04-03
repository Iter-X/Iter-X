import 'package:client/app/constants.dart';
import 'package:client/business/mine/service/profile_service.dart';
import 'package:client/business/mine/widgets/profile_header.dart';
import 'package:client/business/mine/widgets/profile_skeleton.dart';
import 'package:client/business/mine/widgets/section_header.dart';
import 'package:client/business/mine/widgets/stats_card.dart';
import 'package:client/business/mine/widgets/trip_preview_card.dart';
import 'package:client/common/material/app_bar_with_safe_area.dart';
import 'package:client/common/widgets/preference_button.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:provider/provider.dart';

class MinePage extends StatefulWidget {
  const MinePage({super.key});

  @override
  State<MinePage> createState() => _MinePageState();
}

class _MinePageState extends State<MinePage> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      context.read<ProfileService>().fetchUserProfile();
    });
  }

  @override
  Widget build(BuildContext context) {
    return AppBarWithSafeArea(
      backgroundColor: AppColor.bg,
      bottom: false,
      hasAppBar: true,
      actions: [PreferenceButton()],
      child: Consumer<ProfileService>(
        builder: (context, service, child) {
          if (service.isLoading) {
            return const ProfileSkeleton();
          }

          final profile = service.profile;
          if (profile == null) {
            return Center(
                child: Text('Failed to load profile',
                    style: TextStyle(fontSize: 14.sp)));
          }

          return ListView(
            padding: EdgeInsets.only(bottom: 15.h),
            children: [
              ProfileHeader(
                name: profile.userInfo.nickname,
                avatarUrl: profile.userInfo.avatarUrl,
              ),
              SizedBox(height: 20.h),
              Padding(
                padding: EdgeInsets.symmetric(horizontal: 16.w),
                child: Row(
                  children: [
                    Expanded(
                      child: StatsCard(
                        value: profile.exploredCountries.toString(),
                        label: '已探索国家',
                        emoji: '🌍',
                      ),
                    ),
                    SizedBox(width: 10.w),
                    Expanded(
                      child: StatsCard(
                        value: profile.exploredCities.toString(),
                        label: '已探索城市',
                        emoji: '🇨🇳',
                      ),
                    ),
                  ],
                ),
              ),
              SizedBox(height: 10.h),
              Padding(
                padding: EdgeInsets.symmetric(horizontal: 16.w),
                child: Row(
                  children: [
                    Expanded(
                      child: StatsCard(
                        value: profile.exploredStates.toString(),
                        label: '已探索州',
                        emoji: '🇺🇸',
                      ),
                    ),
                    SizedBox(width: 12.w),
                    Expanded(
                      child: StatsCard(
                        value:
                            '${profile.completedBucketListItems}/${profile.totalBucketListItems}',
                        label: '人生清单',
                        emoji: '📝',
                      ),
                    ),
                  ],
                ),
              ),
              SizedBox(height: 20.h),
              Container(
                margin: EdgeInsets.symmetric(horizontal: 15.w),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(AppConfig.boxRadius),
                ),
                child: SectionHeader(
                  title: '人生地图',
                  emoji: '🗺️',
                  onTap: () {
                    // TODO: Navigate to life map page
                  },
                ),
              ),
              SizedBox(height: 20.h),
              Container(
                margin: EdgeInsets.symmetric(horizontal: 15.w),
                decoration: BoxDecoration(
                  color: Colors.white,
                  borderRadius: BorderRadius.circular(AppConfig.boxRadius),
                ),
                child: Column(
                  children: [
                    SectionHeader(
                      title: '行程&游记',
                      emoji: '✨',
                      onTap: () {
                        // TODO: Navigate to trips page
                      },
                    ),
                    ...profile.recentTrips.asMap().entries.map(
                          (entry) => TripPreviewCard(
                            trip: entry.value,
                            isFirst: entry.key == 0,
                            isLast: entry.key == profile.recentTrips.length - 1,
                          ),
                        ),
                  ],
                ),
              ),
            ],
          );
        },
      ),
    );
  }
}
