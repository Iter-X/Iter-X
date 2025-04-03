import 'package:client/app/constants.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';

class SectionHeader extends StatelessWidget {
  final String title;
  final String emoji;
  final VoidCallback? onTap;

  const SectionHeader({
    super.key,
    required this.title,
    required this.emoji,
    this.onTap,
  });

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(AppConfig.boxRadius),
        ),
        child: Padding(
          padding: EdgeInsets.symmetric(horizontal: 15.w, vertical: 10.h),
          child: Row(
            children: [
              Text(
                emoji,
                style: TextStyle(
                  fontSize: 22.sp,
                  fontWeight: AppFontWeight.semiBold,
                ),
              ),
              SizedBox(width: 5.w),
              Text(
                title,
                style: TextStyle(
                  fontSize: 16.sp,
                  fontWeight: AppFontWeight.semiBold,
                  color: AppColor.primaryFont,
                  letterSpacing: 0.5.sp,
                ),
              ),
              const Spacer(),
              Icon(
                Icons.arrow_forward_ios,
                size: 16.sp,
                color: AppColor.grayFont,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
