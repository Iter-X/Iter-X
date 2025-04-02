import 'package:cached_network_image/cached_network_image.dart';
import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class ProfileHeader extends StatelessWidget {
  final String name;
  final String avatarUrl;

  const ProfileHeader({
    super.key,
    required this.name,
    required this.avatarUrl,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: EdgeInsets.symmetric(vertical: 20.h, horizontal: 20.w),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  'Hi, $name 👋',
                  style: TextStyle(
                    fontSize: 30.sp,
                    fontWeight: AppFontWeight.semiBold,
                    color: AppColor.primaryFont,
                  ),
                ),
                SizedBox(height: 8.h),
                Text(
                  'Explore the world',
                  style: TextStyle(
                    fontSize: 18.sp,
                    fontWeight: AppFontWeight.medium,
                    color: AppColor.grayFont,
                    letterSpacing: 0.5.sp,
                  ),
                ),
              ],
            ),
          ),
          Container(
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              border: Border.all(
                color: AppColor.primary,
                width: 1,
              ),
            ),
            child: CircleAvatar(
              radius: 35.w,
              backgroundColor: AppColor.bg,
              child: ClipOval(
                child: CachedNetworkImage(
                  imageUrl: avatarUrl,
                  width: 70.w,
                  height: 70.w,
                  fit: BoxFit.cover,
                  errorWidget: (context, url, error) {
                    return const Icon(Icons.person,
                        size: 40); // TODO: change to a default avatar
                  },
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}
